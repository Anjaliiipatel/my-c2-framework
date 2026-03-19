package main

import (
    "os/exec"
    "runtime"
)

func executeCommand(command string) string {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", command)
    } else {
        cmd = exec.Command("sh", "-c", command)
    }

    out, err := cmd.CombinedOutput()
    if err != nil {
        return err.Error()
    }
    return string(out)
}