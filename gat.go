package main

import (
    "./gat"
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func CommandParser() <-chan string {
    commands := make(chan string, 1)

    go func() {
        for {
            fmt.Print("> ")
            r := bufio.NewReader(os.Stdin)
            in, err := r.ReadString('\n')
            if err != nil {
                log.Fatal(err)
            }

            in = strings.ToLower(strings.TrimSpace(in))
            commands <- in
        }
    }()

    return commands
}

func FileChanged(file string) {
    if filepath.Ext(file) == ".go" {
        fmt.Println("file: ", file)
        fmt.Println("test files: ", gat.TestFiles(file))
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
