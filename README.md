# A simple gRPC service in Golang
## Performs CRUD (Create, Read, Update and Delete) operations:
1. In Golang
2. Using MongoDB <br />
<b>OR</b><br />
Via Docker's MongoDB image <br />
3. (optional) Great tool to [view MongoDB database.](https://robomongo.org)

## Dependencies
1. Install [Golang](https://golang.org/doc/install) and run: <br />
      `$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}` <br />
      `$ go get -u google.golang.org/grpc` <br />
      `$ go get labix.org/v2/mgo` <br />
      
2. Install [MongoDB](https://www.mongodb.com) and run the server: `$ mongod` <br />
<b>OR</b><br />
Install [Docker](https://www.docker.com) and run the following:
`$ docker run -p 27017:27017 -d mongo`

## Cloning the repository
### Clone the repository
`$ git clone https://github.com/maniknarang/gRPC-CRUD-App.git`

## Running the app
Make sure MongoDB server is up and running on `127.0.0.1:27017`. <br />

### If on Mac OS:
`$ cd gRPC-CRUD-App` <br />
`$ chmod 777 scripts/mac_script.sh` <br />
`$ ./scripts/mac_script.sh` <br />
Hit allow on the popup.

### If on Windows/Linux:
Open two terminals and run: `$ cd gRPC-CRUD-App` on both of them. <br /><br />
First terminal: <br />
      `$ go run server/server.go` <br />
      Hit allow on the popup. <br /><br />
Second terminal: <br />
      `$ go run cmd/server/main.go`

## Test File
### Run the test file and see the database to check the changes:
`$ go test tests/server_test.go`
