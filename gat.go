package main

import (
    "fmt"
    "github.com/bobappleyard/readline"
    "github.com/gophertown/gat/gat"
    "github.com/koyachi/go-term-ansicolor/ansicolor"
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

const CSI = "\x1b["

// remove from the screen anything that's been typed
// from github.com/kierdavis/ansi
func ClearPrompt() {
    fmt.Printf("%s2K", CSI)     // clear line
    fmt.Printf("%s%dG", CSI, 0) // go to column 0
}

func GoTest(test_files []string) {
    args := append([]string{"test"}, test_files...)

    cmd := exec.Command("go", args...)

    ClearPrompt()
    fmt.Println(ansicolor.Yellow(strings.Join(cmd.Args, " ")))

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
    }

    fmt.Print(string(out))

    if cmd.ProcessState.Success() {
        fmt.Println(ansicolor.Green("PASS"))
    } else {
        fmt.Println(ansicolor.Red("FAIL"))
    }
}

func Help() {
    fmt.Println(ansicolor.Magenta("\nInteractions:\n"))
    fmt.Println("  * a, all  Run all tests.")
    fmt.Println("  * h, help You found it.")
    fmt.Println("  * e, exit Leave G.A.T.")
}

func Header() {
    fmt.Println(ansicolor.Cyan("G.A.T.0.0.1 is watching your files"))
    fmt.Println("Type " + ansicolor.Magenta("help") + " for help.\n")
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
            ClearPrompt()
            fmt.Println(ansicolor.Yellow("Watching path"), folder)
        case command := <-commands:
            switch command {
            case "exit", "e", "x", "quit", "q":
                break out
            case "all", "a":
                GoTest([]string{"./..."})
            case "help", "h", "?":
                Help()
            default:
                fmt.Println(ansicolor.Red("ERROR:")+" Unknown command", ansicolor.Magenta(command))
            }

        }
    }
}

func main() {
    Header()
    EventLoop()
}
