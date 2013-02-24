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
