package gat

import (
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
