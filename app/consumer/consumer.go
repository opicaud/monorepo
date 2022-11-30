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

func GetRectangleAndSquareArea(address string, request *ac.AreaRequest) (float32, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return 0, err
	}
	defer conn.Close()

	c := ac.NewCalculatorClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalculateArea(ctx, request)

	if err != nil {
		return 0, err
	}
	return r.GetValue(), nil
}
