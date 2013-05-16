package gat

import (
    "log"
    "os/exec"
)

type Run struct {
    Tags string
}

func (run Run) goTest(test_files []string) {
    args := []string{"test"}
    if len(run.Tags) > 0 {
        args = append(args, []string{"-tags", run.Tags}...)
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

func (run Run) RunAll() {
    run.goTest([]string{"./..."})
}

func (run Run) RunOnChange(file string) {
    fc := NewFileChecker(file)
    if fc.IsGoFile() {
        test_files := fc.TestsForGoFile()
        if test_files != nil {
            run.goTest(test_files)
        }
    }
}
