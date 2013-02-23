package main

import (
    "./gat"
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func CommandParser() {
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
}

func main() {
    fmt.Println("G.A.T. 0.0.1 is watching your files")
    watcher := gat.Watch("./")
    CommandParser()
    watcher.Close()
}
