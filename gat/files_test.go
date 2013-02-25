package gat_test

import (
    "./"
    . "launchpad.net/gocheck"
)

func (s *GatSuite) TestTestFile(c *C) {
    c.Assert(gat.TestFile("files.go"), Equals, "files_test.go")
}

func (s *GatSuite) TestSuiteFile(c *C) {
    c.Assert(gat.SuiteFile("folder/files_test.go"), Equals, "folder/suite_test.go")
}

func (s *GatSuite) TestIsSuiteFile(c *C) {
    c.Assert(gat.IsSuiteFile("folder/files_test.go"), Equals, false)
    c.Assert(gat.IsSuiteFile("folder/suite_test.go"), Equals, true)
}

func (s *GatSuite) TestTestForGoFileDoNotExist(c *C) {
    fc := gat.NewFileChecker("files.go")
    fc.Exists = DoNotExist // mock
    c.Assert(fc.TestsForGoFile(), IsNil)
}

func (s *GatSuite) TestTestForGoFileDoExist(c *C) {
    fc := gat.NewFileChecker("files.go")
    fc.Exists = DoExist // mock
    c.Assert(fc.TestsForGoFile(), DeepEquals, []string{"./suite_test.go", "files_test.go"})
}

func (s *GatSuite) TestTestFileDoesNotCheckExistance(c *C) {
    fc := gat.NewFileChecker("files_test.go")
    fc.Exists = DoNotExist // mock
    c.Assert(fc.TestsForGoFile(), DeepEquals, []string{"files_test.go"})
}

// helpers
func DoNotExist(path string) bool {
    return false
}

func DoExist(path string) bool {
    return true
}
