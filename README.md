# Go-C2 Framework 🚀

A lightweight, concurrent Command & Control (C2) framework built with **Go** This project demonstrates how modern malware/agents communicate with a central server using HTTP beaconing.



---

## 🏗️ Architecture

- **Agent (Go):** A cross-platform binary that polls the server for tasks and executes them via the system shell (`cmd` or `sh`).
- **Server (Go + Gin):** A REST API that manages the task queue, stores agent heartbeats, and serves the frontend.
- **Dashboard (React + Vite):** A professional terminal-style UI built with Lovable for monitoring agents and issuing commands.

---

## 🚀 Getting Started

### 1. Prerequisites
- **Go** (1.21+)
- **Node.js & npm**
- **GitHub Codespaces** (or a local terminal)

### 2. Setup & Installation
Clone the repository and install dependencies:

```bash
# Install Go dependencies
go mod tidy

# Install UI dependencies
cd ui
npm install
3. Running the FrameworkYou will need three terminal windows open:Terminal 1: The Backend ServerBashgo run ./server/main.go
Wait for: Listening and serving HTTP on :8080Terminal 2: The AgentBashgo run ./agent/*.go
Wait for: Polling server...Terminal 3: The Frontend (UI)Bashcd ui
npm run dev
🛠️ Commands to TryOnce the agent is "Checking In" (visible in the server logs), go to your Dashboard and type:CommandActionwhoamiCheck the current user contexthostnameGet the name of the target machinels -laList all files (Linux/Mac)dirList all files (Windows)📦 Cross-CompilationOne of Go's superpowers is building the agent for any OS. Run these from the root:Bash# Build for Windows
GOOS=windows GOARCH=amd64 go build -o build/agent.exe ./agent/*.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o build/agent_linux ./agent/*.go
