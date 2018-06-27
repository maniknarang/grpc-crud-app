# Run the server in a new terminal
osascript -e 'tell application "Terminal" to do script "cd ~/go/src/github.com/project; go run server/server.go"'

# Sleep for 5 seconds
sleep 5

# Run the client-side in this terminal
cd ~/go/src/github.com/project
go run cmd/server/main.go

