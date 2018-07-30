# A simple gRPC service in Golang
## Performs CRUD (Create, Read, Update and Delete) operations:
1. In Golang
2. Using MongoDB  
**OR**  
Via Docker's MongoDB image  

## Dependencies
Make sure you have [GOPATH](https://github.com/golang/go/wiki/GOPATH)
environment variable set properly.  
1. Install [Golang](https://golang.org/doc/install) and run:  
      `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`  
      `go get -u google.golang.org/grpc`  
      `go get labix.org/v2/mgo`  
      `go get golang.org/x/net/context`  
      
2. Install [MongoDB](https://www.mongodb.com) and run the server: `mongod`  
**OR**  
Install [Docker](https://www.docker.com) and run the following:
`docker run -p 27017:27017 -d mongo`

## Cloning the repository
### Clone the repository
`cd $HOME/src/github.com/`  
`git clone https://github.com/maniknarang/gRPC-CRUD-App.git`

## Running the app
Make sure MongoDB server is up and running on `127.0.0.1:27017`.  

Open two terminals and run: `cd $HOME/src/github.com/gRPC-CRUD-App` on both of them.    
First terminal:  
      `go run server/server.go`  
      Hit allow on the popup.    
Second terminal:  
      `go run cmd/server/main.go`

## Test File
### Run the test file and see the database to check the changes:
`go run server/server.go`  
`go test tests/server_test.go`