// Autotesting tool with readline support.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	var tags string
	var debug bool
	flag.StringVar(&tags, "tags", "", "a list of build tags for testing.")
	flag.BoolVar(&debug, "debug", false, "adds additional logging")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] [-- [go test options]]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println(`
  EXAMPLE: Specify a -run option to go test, run looper in debug mode:
      looper -debug -- -run MyTest`)
	}
	flag.Parse()

	runner := gat.Run{Tags: tags, Args: flag.Args()}

	Header()
	if debug {
		DebugEnabled()
	}
	EventLoop(runner, debug)
}
