package main

import (
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/project/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	mongo, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}

	defer mongo.Close()

	DB = &mong{mongo.DB("mydb").C("mycol")}

	// Host 127.0.0.1:27017
	listen, err := net.Listen("tcp", ":27017")
	if err != nil {
		log.Fatalf("Could not listen to 27017 port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
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

	var results []interface{}
	e := DB.Operation.Find(bson.M{"id": em.Id}).All(&results)
	if e != nil {
		log.Fatalf("Would never reach here, to make the compiler happy: %v", e)
	}
	if results == nil {
		return nil, status.Error(codes.NotFound, "ID not found")
	}

	// Convert []interface{} to []string for the field "Tags"
	a := (results[0].(bson.M))["tags"].([]interface{})
	b := make([]string, len(a))
	for i := range a {
		b[i] = a[i].(string)
	}

	// Convert map[interface{}]interface{} to map[string]string for the field "Metadata"
	c := (results[0].(bson.M))["metadata"].(bson.M)
	d := make(map[string]string)
	for key, val := range c {
		d[key] = val.(string)
	}

	return &pb.Employee{Name: (results[0].(bson.M))["name"].(string), Id: em.Id,
			Category: int32((results[0].(bson.M))["category"].(int)), Tags: b, Metadata: d},
		DB.Operation.Find(bson.M{"id": em.Id}).One(&em)
}

// UpdateItem updates the item inside the database
// Returns the updated data's ID and error (if any)
func (s *server) UpdateItem(ctx context.Context, em *pb.Employee) (*pb.ID, error) {
	// If ID is null, return specific error
	if em.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	find := bson.M{"id": em.Id}
	update := bson.M{"$set": bson.M{"name": em.Name, "category": em.Category,
		"tags": em.Tags, "metadata": em.Metadata}}
	err := DB.Operation.Update(find, update)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	} else {
		return &pb.ID{Id: em.Id}, nil
	}
}

// DeleteItem deletes the item from the database
// Return the ID of the item deleted and error (if any)
func (s *server) DeleteItem(ctx context.Context, em *pb.ID) (*pb.ID, error) {
	return &pb.ID{Id: em.Id}, nil
}
