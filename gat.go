package main

import (
    "github.com/bobappleyard/readline"
    "github.com/gophertown/gat/gat"
    "io"
    "log"
    "strings"
)

func CommandParser() <-chan string {
    commands := make(chan string, 1)

    go func() {
        for {
            in, err := readline.String("")
            if err == io.EOF { // Ctrl+D
                commands <- "eof"
                break
            } else if err != nil {
                log.Fatal(err)
            }

            in = strings.ToLower(strings.TrimSpace(in))
            commands <- in
            readline.AddHistory(in)
        }
    }()

    return commands
}

func FileChanged(file string) {
    if gat.IsGoFile(file) {
        test_files := gat.TestsForGoFile(file)
        if test_files != nil {
            gat.GoTest(test_files)
        }
    }
}

func EventLoop() {
    commands := CommandParser()
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
            case "exit", "e", "x", "quit", "q", "eof":
                break out
            case "all", "a":
                gat.GoTestAll()
            case "help", "h", "?":
                gat.Help()
            default:
                gat.UnknownCommand(command)
            }

        }
    }
}

func main() {
    gat.Header()
    EventLoop()
}
