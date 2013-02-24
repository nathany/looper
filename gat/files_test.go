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
}

func (s *GatSuite) TestTestFiles(c *C) {
    for _, t := range fileTests {
        c.Check(gat.TestFiles(t.in), DeepEquals, t.out)
    }
}
