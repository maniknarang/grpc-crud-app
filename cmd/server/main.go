package main

import (
	"log"
	"os"
	"time"

	pb "github.com/project/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Change the host here only
const (
	address     = "localhost:27017"
	defaultName = "Manik"
)

func main() {
	// Connect to the server
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

	// Menu-based program allowing the user to choose from CRUD

	// CreateItem operation
	item, err := c.CreateItem(ctx, &pb.Employee{Name: name, Id: "1234"})
	if err != nil {
		log.Fatalf("Could not create a new item: %v", err)
	}
	log.Printf("Name: %s", item.Id)

	// ReadItem operation
	read, err := c.ReadItem(ctx, &pb.ID{Id: "1234"})
	if err != nil {
		log.Fatalf("Error reading the item: %v", err)
	}
	log.Printf("Item found: %s", read.Name)

	// UpdateItem operation
	up, err := c.UpdateItem(ctx, &pb.Employee{Name: "LOL1", Id: "1234"})
	if err != nil {
		log.Fatalf("Error updating the item: %v", err)
	}
	log.Printf("Item updated with the ID: %s", up.Id)

	// DeleteItem operation
	// Ignoring the error - should always be successful regardless of the
	// implicit find result
	del, _ := c.DeleteItem(ctx, &pb.ID{Id: "1234"})
	log.Printf("Item with the ID %s deleted", del.Id)
}
