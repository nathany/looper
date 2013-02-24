package gat

import (
    "log"
    "os/exec"
)

func GoTest(test_files []string) {
    args := append([]string{"test"}, test_files...)

    cmd := exec.Command("go", args...)

    PrintCommand(cmd.Args) // includes "go"

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
    }
    PrintCommandOutput(out)

    RedGreen(cmd.ProcessState.Success())
}

func GoTestAll() {
    GoTest([]string{"./..."})
}
