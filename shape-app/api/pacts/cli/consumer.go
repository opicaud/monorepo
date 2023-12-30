package cli

import (
	"context"
	"github.com/opicaud/monorepo/shape-app/api/proto"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Client struct{}
type Features interface {
	CreateShape(address string, request *proto.ShapeRequest) (*proto.Message, error)
}

func (f Client) CreateShape(address string, request *proto.ShapeRequest) (*proto.Message, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := proto.NewShapesClient(conn)

	logger.Info("Sending", "shapeRequest", request)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}
