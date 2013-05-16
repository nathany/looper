// Autotesting tool with readline support.
package main

import (
    "flag"
    "github.com/gophertown/looper/gat"
    "log"
)

type Args struct {
    Tags string
}

func EventLoop(args *Args) {
    commands := gat.CommandParser()
    watcher, err := gat.NewRecurisveWatcher("./")
    if err != nil {
        log.Fatal(err)
    }
    watcher.Run()
    defer watcher.Close()

out:
    for {
        select {
        case file := <-watcher.Files:
            FileChanged(args.Tags, file)
        case folder := <-watcher.Folders:
            gat.PrintWatching(folder)
        case command := <-commands:
            switch command {
            case gat.EXIT:
                break out
            case gat.TEST_ALL:
                gat.GoTestAll(args.Tags)
            case gat.HELP:
                gat.Help()
            }
        }
    }
}

func FileChanged(tags string, file string) {
    fc := gat.NewFileChecker(file)
    if fc.IsGoFile() {
        test_files := fc.TestsForGoFile()
        if test_files != nil {
            gat.GoTest(tags, test_files)
        }
    }
}

func main() {
    var args Args
    flag.StringVar(&args.Tags, "tags", "", "a list of build tags for testing.")
    flag.Parse()
    gat.Header()
    EventLoop(&args)
}
