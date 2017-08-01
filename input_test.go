package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/koyachi/go-term-ansicolor/ansicolor"
)

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return fmt.Sprintf("can't create pipe: %s", err.Error())
	}
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNormalizeCommand(t *testing.T) {
	sudo := "sudo"
	commands := []struct {
		input string
		cmd   Command
	}{
		{" Exit", Exit},
		{sudo, Unknown},
	}

	output := captureOutput(func() {
		for _, c := range commands {
			actual := NormalizeCommand(c.input)
			if actual != c.cmd {
				t.Errorf("Expected '%s' to result in %v, but got %v", c.input, c.cmd, actual)
			}
		}
	})

	expected := fmt.Sprintln(ansicolor.Red("ERROR:")+" Unknown command", ansicolor.Magenta(sudo))

	if output != expected {
		t.Errorf("Expected output '%s', but got '%s'", output, expected)
	}
}
