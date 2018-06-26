package main

import (
	"log"
    "os"
    "time"

//	"gopkg.in/mgo.v2"

	pb "github.com/project/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:27017"
    defaultName = "Manik"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Did not connect/dial: %v", err)
    }
    defer conn.Close()
    c := pb.NewCRUDClient(conn)

    name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()
    r, err := c.CreateItem(ctx, &pb.Employee{Name: name, Id: "1234"})
    if err != nil {
        log.Fatalf("could not insert: %v", err)
    }
    log.Printf("Name: ", r.Id)
}
