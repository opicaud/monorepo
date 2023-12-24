package pacts

import (
	"context"
	"github.com/opicaud/monorepo/shape-app/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Foo struct{}
type Features interface {
	GetRectangleAndSquareArea2(address string, request *proto.ShapeRequest) (*proto.Message, error)
}

func (f Foo) GetRectangleAndSquareArea2(address string, request *proto.ShapeRequest) (*proto.Message, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := proto.NewShapesClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}
