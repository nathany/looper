package main

import (
    "github.com/bobappleyard/readline"
    "github.com/gophertown/gat/gat"
    "io"
    "log"
    "os/exec"
    "strings"
)

func CommandParser() <-chan string {
    commands := make(chan string, 1)

    go func() {
        for {
            in, err := readline.String("")
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
        test_files := gat.TestsForGoFile(file)
        if test_files != nil {
            GoTest(test_files)
        }
    }
}

func GoTest(test_files []string) {
    args := append([]string{"test"}, test_files...)

    cmd := exec.Command("go", args...)

    gat.PrintCommand(cmd.Args) // includes "go"

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
    }
    gat.PrintCommandOutput(out)

    gat.RedGreen(cmd.ProcessState.Success())
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
            case "exit", "e", "x", "quit", "q":
                break out
            case "all", "a":
                GoTest([]string{"./..."})
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
