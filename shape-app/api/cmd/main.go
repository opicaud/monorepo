package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/opicaud/monorepo/shape-app/api/proto"
	"github.com/opicaud/monorepo/shape-app/domain/shape"
	"github.com/opicaud/monorepo/shape-app/events/eventstore/inmemory/cmd"
	"github.com/opicaud/monorepo/shape-app/events/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedShapesServer
}

func (s *server) Create(ctx context.Context, in *pb.ShapeRequest) (*pb.Response, error) {
	repository := cmd.NewInMemoryEventStore()
	factory := shape.NewFactory()
	var command = factory.NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions...)

	provider := pkg.NewEventsFrameworkBuilder().WithEventStore(repository).Build()
	handler := shape.NewShapeCreationCommandHandlerBuilder().WithInfraProvider(provider).Build()
	err := handler.Execute(command)
	if err != nil {
		return nil, err
	}
	message := pb.Message{
		Code: uint32(codes.OK),
	}
	response := pb.Response{
		Message: &message,
	}
	return &response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShapesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
