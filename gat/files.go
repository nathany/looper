package gat

import (
    "log"
    "os"
    "regexp"
)

var (
    isTestFile         = regexp.MustCompile(`_test\.go$`)
    architectureSuffix = regexp.MustCompile(`(.+)_(386|amd64|arm)$`)
    osSuffix           = regexp.MustCompile(`(.+)_(darwin|freebsd|linux|netbsd|openbsd|plan9|windows|unix)$`)
)

func TestFiles(file string) []string {
    if isTestFile.MatchString(file) {
        return []string{file}
    }

    file = file[:len(file)-3]
    files := []string{file + "_test.go"}

    matches := architectureSuffix.FindStringSubmatch(file)
    if len(matches) > 0 {
        file = matches[1]
        files = append(files, file+"_test.go")
    }

    matches = osSuffix.FindStringSubmatch(file)
    if len(matches) > 0 {
        file = matches[1]
        files = append(files, file+"_test.go")
    }

    return files
}

// FIXME: Stat on the file system is probably a bit inefficient
func Exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if !os.IsNotExist(err) {
        log.Fatal(err)
    }
    return false
}

func Filter(vs []string, f func(string) bool) (filtered []string) {
    for _, s := range vs {
        if f(s) {
            filtered = append(filtered, s)
        }
    }
    return filtered
}

// FIXME: only need filter when deriving file names (not for change to _test file)
func TestFilesThatExist(file string) []string {
    return Filter(TestFiles(file), Exists)
}
