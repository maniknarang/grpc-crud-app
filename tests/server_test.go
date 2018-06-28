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

// Test for CreateItem method when input has empty name
func Test_Integration_CreateItem_EmptyName(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "", Id: "123", Category: 123,
		Tags:     []string{"this", "is", "an", "array", "of", "strings"},
		Metadata: map[string]string{"hello": "world"}})

	// No error as insertion succeeds
	assert.Nil(t, err)
	// ID's should be equal as insertion succeeds
	assert.Equal(t, a.Id, "123", "Insertion should be fine, so they should",
		"equal each other")
	// ID's should not be equal here
	assert.NotEqual(t, a.Id, "00000", "ID's are not equal")
}

// Test for CreateItem method when input has empty ID
func Test_Integration_CreateItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "", Id: "", Category: 321,
		Tags:     []string{"this", "is", "an", "array", "of", "strings"},
		Metadata: map[string]string{"hello": "world"}})

	// Error as insertion fails - ID is empty
	assert.NotNil(t, err)
	// a should be nil as insertion fails
	assert.Nil(t, a)
}

// Test for CreateItem method when input is normal
func Test_Integration_CreateItem_NormalInput(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.CreateItem(ctx, &pb.Employee{Name: "Heyo", Id: "12345", Category: 132,
		Tags:     []string{"this", "is", "an", "array", "of", "strings"},
		Metadata: map[string]string{"hello": "world"}})

	// No error as insertion succeeds
	assert.Nil(t, err)
	// ID's should be equal as insertion succeeds
	assert.Equal(t, a.Id, "12345", "Insertion should be fine, so they should",
		"equal each other")
	// ID's should not be equal here
	assert.NotEqual(t, a.Id, "00000", "ID's are not equal")
}

// Test for ReadItem method when input ID is empty
func Test_Integration_ReadItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.ReadItem(ctx, &pb.ID{Id: ""})

	// Error occurs as ID is empty
	assert.NotNil(t, err)
	// Search results are nil as ID is empty
	assert.Nil(t, a)
}

// Test for ReadItem method when input ID is normal
func Test_Integration_ReadItem_NormalID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.ReadItem(ctx, &pb.ID{Id: "12345"})

	// Error should be nil as find passed
	assert.Nil(t, err)
	// ID's should be equal as find passed
	assert.Equal(t, a.Id, "12345", "ID's are equal")
	// Names should be equal as find passed
	assert.Equal(t, a.Name, "Heyo", "Names are equal")
	// Categories should be equal as find passed
	assert.Equal(t, a.Category, int32(132), "Categories are equal")
	// Tags should be equal as find passed
	assert.Equal(t, a.Tags, []string{"this", "is", "an", "array", "of", "strings"}, "Tags are equal")
	// Metadata should be equal as find passed
	assert.Equal(t, a.Metadata, map[string]string{"hello": "world"}, "Metadata are equal")
}

// Test for UpdateItem method when input has empty ID
func Test_Integration_UpdateItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.UpdateItem(ctx, &pb.Employee{Id: ""})

	// Update fails so error is not nil
	assert.NotNil(t, err)
	// Update fails so a is nil
	assert.Nil(t, a)
}

// Test for UpdateItem method when input is all normal
func Test_Integration_UpdateItem_NormalInput(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.UpdateItem(ctx, &pb.Employee{Name: "NowHasAName", Id: "123", Category: 321,
		Tags: []string{"updated", "tags"}, Metadata: map[string]string{"heyo": "world"}})

	// Error should be nil as the update passed
	assert.Nil(t, err)
	// ID's should be equal as the update passed
	assert.Equal(t, a.Id, "123", "ID's are equal")
}

// Test for DeleteItem method when input has empty ID
func Test_Integration_DeleteItem_EmptyID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.DeleteItem(ctx, &pb.ID{Id: ""})

	// Delete is always successful so no error
	assert.Nil(t, err)
	// Delete is always successful so a exists with zeroed out fields
	assert.NotNil(t, a)
}

// Test for DeleteItem method when input is normal
func Test_Integration_DeleteItem_NormalID(t *testing.T) {
	conn, err := grpc.Dial("localhost:27017", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	a, err := c.DeleteItem(ctx, &pb.ID{Id: "123"})

	// Delete is always successful so no error
	assert.Nil(t, err)
	// Delete is always successful so a exists with zeroed out fields
	assert.NotNil(t, a)
}
