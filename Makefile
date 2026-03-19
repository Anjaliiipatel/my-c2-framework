build-agent-windows:
	GOOS=windows GOARCH=amd64 go build -o build/agent.exe ./agent/*.go

build-agent-linux:
	GOOS=linux GOARCH=amd64 go build -o build/agent-linux ./agent/*.go

run-server:
	go run server/main.go