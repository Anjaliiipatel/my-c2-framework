package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // CRITICAL: Allow Lovable frontend to talk to this API
    r.Use(cors.Default())

    // Agent Endpoints
    r.GET("/api/v1/poll/:agent_id", GetPendingTask)
    r.POST("/api/v1/results", PostTaskResult)

    // UI Endpoints (for Lovable)
    r.POST("/api/v1/admin/command", CreateTask)
    r.GET("/api/v1/admin/agents", ListAgents)

    r.Run(":8080")
}