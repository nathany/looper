package main

import (
	. "gopkg.in/check.v1"
)

func (s *LooperSuite) TestMixedCaseCommand(c *C) {
	c.Assert(NormalizeCommand(" Exit"), Equals, EXIT)
}

func (s *LooperSuite) TestUnkownCommand(c *C) {
	c.Assert(NormalizeCommand("sudo"), Equals, UNKNOWN)
}
