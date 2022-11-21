package consumer

import (
	"context"
	ac "example2/proto"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func GetRectangleAndSquareArea(address string) ([]float32, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := ac.NewCalculatorClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rectangle := ac.ShapeMessage{Shape: &ac.ShapeMessage_Rectangle{Rectangle: &ac.Rectangle{Length: 3, Width: 4}}}
	square := ac.ShapeMessage{Shape: &ac.ShapeMessage_Square{Square: &ac.Square{EdgeLength: 3}}}
	r, err := c.CalculateMulti(ctx, &ac.AreaRequest{Shapes: []*ac.ShapeMessage{&rectangle, &square}})
	if err != nil {
		return nil, err
	}
	return r.GetValue(), nil
}
