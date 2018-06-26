# Run the client-side in a different terminal
osascript -e 'tell application "Terminal" to do script "cd ~/go/src/github.com/project; go run server/server.go"'

# Sleep for 5 seconds

# Run the server in this terminal
cd ~/go/src/github.com/project
go run cmd/server/main.go

