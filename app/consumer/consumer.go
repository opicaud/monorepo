package consumer

import (
	"context"
	ac "example2/app/proto"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func GetRectangleAndSquareArea(address string, request *ac.ShapeRequest) (*ac.Message, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := ac.NewShapesClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}
