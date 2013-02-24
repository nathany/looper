package main

import (
    "fmt"
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
            in, err := readline.String("> ")
            if err == io.EOF { // Ctrl+D
                commands <- "exit"
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
        fmt.Println("file: ", file)
        fmt.Println("test files: ", gat.TestsForGoFile(file))
    }
}

func main() {
    fmt.Println("G.A.T. 0.0.1 is watching your files")

    watcher := gat.Watch("./")
    commands := CommandParser()

out:
    for {
        select {
        case file := <-watcher.Files:
            FileChanged(file)
        case command := <-commands:
            if command == "exit" {
                break out
            }
            fmt.Println("command: ", command)
        }
    }
    watcher.Close()
}
