package models

type Task struct {
    ID      string `json:"id"`
    Command string `json:"command"`
    Result  string `json:"result"`
    Status  string `json:"status"` // "pending" or "completed"
}

type Agent struct {
    ID       string `json:"id"`
    Hostname string `json:"hostname"`
    Platform string `json:"platform"`
}