package main

import (
    "bufio"
    "fmt"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
    "strings"
)

func watch() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case event := <-watcher.Event:
                fmt.Print("\n", event, "\n> ")
            case err := <-watcher.Error:
                log.Println("error", err)
            }
        }
    }()

    err = watcher.Watch("./")
    if err != nil {
        log.Fatal(err)
    }

    err = watcher.Watch("./b")
    if err != nil {
        log.Fatal(err)
    }

    // do stuff
    for {
        fmt.Print("> ")
        r := bufio.NewReader(os.Stdin)
        in, err := r.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        in = strings.ToLower(strings.TrimSpace(in))
        if in == "exit" {
            break
        }

        fmt.Println(in)
    }

    watcher.Close()
}

func main() {
    fmt.Println("G.A.T. 0.0.1 is watching your files")
    watch()
}
