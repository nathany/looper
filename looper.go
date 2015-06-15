// Autotesting tool with readline support.
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nathany/looper/gat"
)

type Runner interface {
	RunOnChange(file string)
	RunAll()
}

func EventLoop(runner Runner, debug bool) {
	commands := CommandParser()
	watcher, err := NewRecursiveWatcher("./")
	if err != nil {
		log.Fatal(err)
	}
	watcher.Run(debug)
	defer watcher.Close()

out:
	for {
		select {
		case file := <-watcher.Files:
			runner.RunOnChange(file)
		case folder := <-watcher.Folders:
			PrintWatching(folder)
		case command := <-commands:
			switch command {
			case Exit:
				break out
			case RunAll:
				runner.RunAll()
			case Help:
				DisplayHelp()
			}
		}
	}
}

func main() {
	// Get debug status from env var, if error ignore and debug is off
	debug, _ := strconv.ParseBool(os.Getenv("LOOPER_DEBUG"))

	// Pass all args to go test, except the name of the looper command
	gtargs := os.Args[1:len(os.Args)]
	runner := gat.Run{Args: gtargs}

	Header(gtargs)
	if debug {
		DebugEnabled()
	}
	EventLoop(runner, debug)
}
