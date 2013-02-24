package gat_test

import (
    "./"
    . "launchpad.net/gocheck"
)

var fileTests = []struct {
    in  string
    out []string
}{
    {"files_test.go", []string{"files_test.go"}},
    {"files.go", []string{"files_test.go"}},
    {"files_darwin.go", []string{"files_darwin_test.go", "files_test.go"}},
    {"files_amd64.go", []string{"files_amd64_test.go", "files_test.go"}},
    {"files_windows_386.go", []string{"files_windows_386_test.go", "files_windows_test.go", "files_test.go"}},
}

func (s *GatSuite) TestTestFiles(c *C) {
    for _, t := range fileTests {
        c.Check(gat.TestFiles(t.in), DeepEquals, t.out)
    }
}
