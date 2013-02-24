package gat

import (
    "regexp"
)

var (
    isTestFile = regexp.MustCompile(`_test\.go$`)
)

func TestFiles(file string) []string {
    if isTestFile.MatchString(file) {
        return []string{file}
    }
    file = file[:len(file)-3] + "_test.go"
    return []string{file}
}
