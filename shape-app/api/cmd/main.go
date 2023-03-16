package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/opicaud/monorepo/events/eventstore/inmemory/cmd"
	"github.com/opicaud/monorepo/events/pkg"
	pb "github.com/opicaud/monorepo/shape-app/api/proto"
	"github.com/opicaud/monorepo/shape-app/domain/shape"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedShapesServer
}

func (s *server) Create(ctx context.Context, in *pb.ShapeRequest) (*pb.Response, error) {
	factory := shape.NewFactory()
	var command = factory.NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions...)

	eventsFramework := pkg.NewEventsFrameworkBuilder().WithEventStore(cmd.NewInMemoryEventStore()).Build()
	handler := shape.NewCommandHandlerBuilder().WithEventsFramework(eventsFramework).Build()
	err := handler.Execute(command, shape.NewShapeCommandApplier(eventsFramework))
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
