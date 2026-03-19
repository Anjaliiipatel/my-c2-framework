package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "my-c2/server/models" // Adjust based on your module name
)

func fetchTask(url string) (*models.Task, error) {
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != http.StatusOK {
        return nil, err
    }
    defer resp.Body.Close()

    var t models.Task
    err = json.NewDecoder(resp.Body).Decode(&t)
    return &t, err
}

func postResult(url string, task models.Task) error {
    data, _ := json.Marshal(task)
    _, err := http.Post(url, "application/json", bytes.NewBuffer(data))
    return err
}