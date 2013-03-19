package gat

import (
    "log"
    "os/exec"
)

func GoTest(tags string, test_files []string) {
    args := []string{"test"}
    if len(tags) > 0 {
        args = append(args, []string{"-tags", tags}...)
    }
    args = append(args, test_files...)

    cmd := exec.Command("go", args...)
    // cmd.Dir watchDir = ./

    PrintCommand(cmd.Args) // includes "go"

    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Println(err)
    }
    PrintCommandOutput(out)

    RedGreen(cmd.ProcessState.Success())
    ShowDuration(cmd.ProcessState.UserTime())
}

func GoTestAll(tags string) {
    GoTest(tags, []string{"./..."})
}
