package gat

import (
    "fmt"
    "github.com/koyachi/go-term-ansicolor/ansicolor"
    "strings"
)

func Header() {
    fmt.Println(ansicolor.Cyan("G.A.T.0.0.1 is watching your files"))
    fmt.Println("Type " + ansicolor.Magenta("help") + " for help.\n")
}

func Help() {
    fmt.Println(ansicolor.Magenta("\nInteractions:\n"))
    fmt.Println("  * a, all  Run all tests.")
    fmt.Println("  * h, help You found it.")
    fmt.Println("  * e, exit Leave G.A.T.")
}

func PrintCommand(args []string) {
    ClearPrompt()
    fmt.Println(ansicolor.Yellow(strings.Join(args, " ")))
}

func PrintCommandOutput(out []byte) {
    fmt.Print(string(out))
}

func RedGreen(pass bool) {
    if pass {
        fmt.Println(ansicolor.Green("PASS"))
    } else {
        fmt.Println(ansicolor.Red("FAIL"))
    }
}

func PrintWatching(folder string) {
    ClearPrompt()
    fmt.Println(ansicolor.Yellow("Watching path"), ansicolor.Yellow(folder))
}

func UnknownCommand(command string) {
    fmt.Println(ansicolor.Red("ERROR:")+" Unknown command", ansicolor.Magenta(command))
}

const CSI = "\x1b["

// remove from the screen anything that's been typed
// from github.com/kierdavis/ansi
func ClearPrompt() {
    fmt.Printf("%s2K", CSI)     // clear line
    fmt.Printf("%s%dG", CSI, 0) // go to column 0
}
