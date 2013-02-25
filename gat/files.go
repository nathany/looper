package gat

import (
    "log"
    "os"
    "path/filepath"
    "regexp"
)

var (
    testFileMatcher = regexp.MustCompile(`_test\.go$`)
)

type FileChecker struct {
    file   string
    Exists func(string) bool
}

func NewFileChecker(file string) FileChecker {
    return FileChecker{file: file, Exists: FileExists}
}

func (fc FileChecker) IsGoFile() bool {
    return filepath.Ext(fc.file) == ".go"
}

func (fc FileChecker) TestsForGoFile() []string {
    // if the suite file triggers a change, run tests against the entire folder
    if IsSuiteFile(fc.file) {
        return []string{"./" + filepath.Dir(fc.file)} // watchDir = ./
    }

    var test_file string
    // test file triggered a modify/create, we know it exists
    if fc.IsTestFile() {
        test_file = fc.file
    } else {
        test_file = TestFile(fc.file)
        // no tests to run
        if !fc.Exists(test_file) {
            return nil
        }
    }

    suite_file := SuiteFile(test_file)
    // if not found here, should it look in the parent folder?
    if !fc.Exists(suite_file) {
        return []string{test_file}
    }

    return []string{suite_file, test_file}
}

func (fc FileChecker) IsTestFile() bool {
    return testFileMatcher.MatchString(fc.file)
}

func TestFile(file string) string {
    file = file[:len(file)-3] + "_test.go"
    return file
}

func SuiteFile(file string) string {
    return filepath.Dir(file) + "/suite_test.go"
}

func IsSuiteFile(file string) bool {
    return filepath.Base(file) == "suite_test.go"
}

func FileExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if !os.IsNotExist(err) {
        log.Fatal(err)
    }
    return false
}

// returns a slice of subfolders (recursive), including the folder passed in
func Subfolders(path string) (paths []string) {
    filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            name := info.Name()
            // skip folders that begin with a dot
            hidden := filepath.HasPrefix(name, ".") && name != "." && name != ".."
            if hidden {
                return filepath.SkipDir
            } else {
                paths = append(paths, newPath)
            }
        }
        return nil
    })
    return paths
}
