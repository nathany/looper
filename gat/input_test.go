package gat_test

import (
    "./"
    . "launchpad.net/gocheck"
)

func (s *GatSuite) TestMixedCaseCommand(c *C) {
    c.Assert(gat.NormalizeCommand(" Exit"), Equals, gat.EXIT)
}

func (s *GatSuite) TestUnkownCommand(c *C) {
    c.Assert(gat.NormalizeCommand("sudo"), Equals, gat.UNKNOWN)
}
