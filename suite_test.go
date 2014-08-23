package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type LooperSuite struct{}

var _ = Suite(&LooperSuite{})
