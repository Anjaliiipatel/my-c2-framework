# my-c2-framework
my-c2-framework/
├── agent/               # The Go code for the target binary
│   ├── main.go          # Entry point (infinite loop)
│   ├── execute.go       # Logic for os/exec
│   └── client.go        # HTTP client & heartbeat logic
├── server/              # The Go backend (REST API)
│   ├── main.go          # Gin or Echo server setup
│   ├── handlers/        # API routes (e.g., GetTasks, PostResults)
│   ├── models/          # Shared structs (Task, Agent, Command)
│   └── store/           # Database logic (SQLite or PostgreSQL)
├── ui/                  # The Lovable/React frontend
│   ├── src/             # Dashboard, Terminal components
│   └── ...              # Vite/React config
├── pkg/                 # Shared Go packages
│   └── crypto/          # Shared encryption helpers
└── Makefile             # Shortcuts to build for different OS/Arch