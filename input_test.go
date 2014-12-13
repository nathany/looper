package main

import (
	"testing"
)

func TestNormalizeCommand(t *testing.T) {
	commands := []struct {
		input string
		cmd   Command
	}{
		{" Exit", EXIT},
		{"sudo", UNKNOWN},
	}

	for _, c := range commands {
		actual := NormalizeCommand(c.input)
		if actual != c.cmd {
			t.Errorf("Expected '%s' to result in %v, but got %v", c.input, c.cmd, actual)
		}
	}
}
