package main

import (
	"log"
	"net"

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

// DB is a pointer to the mongo struct (using mango)
var DB *mong

func main() {
	mongo, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}

	defer mongo.Close()

	DB = &mong{mongo.DB("mydb").C("mycol")}

	// gRPC server
	listen, err := net.Listen("tcp", ":27017")
	if err != nil {
		log.Fatalf("Could not listen to 27017 port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) CreateItem(ctx context.Context, em *pb.Employee) (*pb.ID, error) {
	return &pb.ID{Id: "1234"}, DB.Operation.Insert(em)
}

func (s *server) ReadItem(ctx context.Context, em *pb.ID) (*pb.Employee, error) {
	return &pb.Employee{}, DB.Operation.Find(bson.M{"id": "1234"}).One(&em)
}

func (s *server) UpdateItem(ctx context.Context, em *pb.Employee) (*pb.ID, error) {
	return &pb.ID{Id: em.Id}, DB.Operation.Update(bson.M{"$set": bson.M{"name": em.Name}}, bson.M{"id": em.Id})
}

func (s *server) DeleteItem(ctx context.Context, em *pb.ID) (*pb.ID, error) {
	return &pb.ID{Id: em.Id}, DB.Operation.Remove(bson.M{"id": em.Id})
}
