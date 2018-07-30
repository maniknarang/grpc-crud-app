package main

import (
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gRPC-CRUD-App/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type server struct{}

type mong struct {
	Operation *mgo.Collection
}

// DB is a pointer to the mong struct (using mango/mgo)
var DB *mong

func main() {
	// Host mongo service
	mongo, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer mongo.Close()

	DB = &mong{mongo.DB("mydb").C("mycol")}

	// Host grpc service
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}

// CreateItem creates a new item in the database
// Returns the inserted ID and error (if any)
func (s *server) CreateItem(ctx context.Context, em *pb.Employee) (*pb.ID, error) {
	// If ID is null, return specific error
	if em.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	return &pb.ID{Id: em.Id}, DB.Operation.Insert(em)
}

// ReadItem reads an item using the ID in the database
// Returns the employee name and ID and error (if any)
func (s *server) ReadItem(ctx context.Context, em *pb.ID) (*pb.Employee, error) {
	// If ID is null, return specific error
	if em.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	var result pb.Employee
	err := DB.Operation.Find(bson.M{"id": em.Id}).One(&result)
	if err != nil {
		log.Printf("Error retrieving employee with id: %s, error: %v", em.Id, err)
		return nil, err
	}

	return &result, nil
}

// UpdateItem updates the item inside the database
// Returns the updated data's ID and error (if any)
func (s *server) UpdateItem(ctx context.Context, em *pb.Employee) (*pb.ID, error) {
	// If ID is null, return specific error
	if em.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	find := bson.M{"id": em.Id}
	update := bson.M{"$set": bson.M{"name": em.Name, "category": em.Category, "tags": em.Tags, "metadata": em.Metadata}}

	return &pb.ID{Id: em.Id}, DB.Operation.Update(find, update)
}

// DeleteItem deletes the item from the database
// Return the ID of the item deleted and error (if any)
func (s *server) DeleteItem(ctx context.Context, em *pb.ID) (*pb.ID, error) {
	return &pb.ID{Id: em.Id}, DB.Operation.Remove(bson.M{"id": em.Id})
}
