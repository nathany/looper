package main

import (
    "github.com/bobappleyard/readline"
    "github.com/gophertown/gat/gat"
    "github.com/kierdavis/ansi"
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

    ansi.ClearLine()
    ansi.CursorHozPosition(0)
    ansi.Println(ansi.Yellow, strings.Join(cmd.Args, " "))

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
    }

    ansi.Print(ansi.White, string(out))

    if cmd.ProcessState.Success() {
        ansi.Println(ansi.Green, "PASS")
    } else {
        ansi.Println(ansi.Red, "FAIL")
    }
}

func Help() {
    ansi.Println(ansi.Magenta, "\nHelp:\n")
    ansi.Println(ansi.White, "  * a, all  Run all tests.")
    ansi.Println(ansi.White, "  * h, help You found it.")
    ansi.Println(ansi.White, "  * e, exit Leave G.A.T.")

}

func main() {
    ansi.Println(ansi.Cyan, "G.A.T.0.0.1 is watching your files")
    ansi.Print(ansi.White, "Type ")
    ansi.Print(ansi.Magenta, "help ")
    ansi.Println(ansi.White, "for help.\n")

    watcher := gat.Watch("./")
    commands := CommandParser()

out:
    for {
        select {
        case file := <-watcher.Files:
            FileChanged(file)
        case command := <-commands:
            switch command {
            case "exit", "e", "x", "quit", "q":
                break out
            case "all", "a":
                GoTest([]string{"./..."})
            case "help", "h", "?":
                Help()
            default:
                ansi.Print(ansi.Red, "ERROR: ")
                ansi.Println(ansi.White, "Unknown command", command)
            }

        }
    }
    watcher.Close()
}
