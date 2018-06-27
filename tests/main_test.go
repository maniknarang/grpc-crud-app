package grpc

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/project/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func Test_Integration_CreateItem_EmptyName(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "", Id: "123"})

	assert.Nil(t, err)
	assert.Equal(t, a.Id, "123", "Insertion should be fine, so they should",
		"equal each other")
	assert.NotEqual(t, a.Id, "00000", "ID's are not equal")
}

func Test_Integration_CreateItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "", Id: ""})

	assert.Nil(t, err)
	assert.Equal(t, a.Id, "", "Insertion should be fine, so they should",
		"equal each other")
	assert.NotEqual(t, a.Id, "00000", "ID's are not equal")
}

func Test_Integration_CreateItem_NormalInput(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "Heyo", Id: "12345"})

	assert.Nil(t, err)
	assert.Equal(t, a.Id, "12345", "Insertion should be fine, so they should",
		"equal each other")
	assert.NotEqual(t, a.Id, "00000", "ID's are not equal")
}

/*func Test_Integration_ReadItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.ReadItem(ctx, &pb.ID{Id: ""})

	assert.NotNil(t, err)
	assert.Nil(t, a)
} */

func Test_Integration_ReadItem_NormalID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "Heyo", Id: "12345"})
	b, err := c.ReadItem(ctx, &pb.ID{Id: "12345"})

	assert.Nil(t, err)
	assert.Equal(t, a.Id, "12345", "ID's are equal")
	assert.Equal(t, b.Id, "12345", "ID's are equal")
	assert.Equal(t, b.Name, "Heyo", "ID's are equal")
	assert.NotEqual(t, b.Id, "00000", "ID's are not equal")
}
