package main

import (
    "github.com/gophertown/gat/gat"
    "log"
)

func EventLoop() {
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
            FileChanged(file)
        case folder := <-watcher.Folders:
            gat.PrintWatching(folder)
        case command := <-commands:
            switch command {
            case gat.EXIT:
                break out
            case gat.TEST_ALL:
                gat.GoTestAll()
            case gat.HELP:
                gat.Help()
            }
        }
    }
}

func FileChanged(file string) {
    fc := gat.NewFileChecker(file)
    if fc.IsGoFile() {
        test_files := fc.TestsForGoFile()
        if test_files != nil {
            gat.GoTest(test_files)
        }
    }
}

func main() {
    gat.Header()
    EventLoop()
}
