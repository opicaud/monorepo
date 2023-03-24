package main

import (
	"context"
	"flag"
	"fmt"
	pkg3 "github.com/opicaud/monorepo/events/eventstore/pkg"
	"github.com/opicaud/monorepo/events/pkg"
	pb "github.com/opicaud/monorepo/shape-app/api/proto"
	pkg2 "github.com/opicaud/monorepo/shape-app/domain/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedShapesServer
}

func (s *server) Create(ctx context.Context, in *pb.ShapeRequest) (*pb.Response, error) {
	factory := pkg2.New()
	var command = factory.NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions...)
	eventStore, errConfig := pkg3.NewEventsFrameworkFromConfig(os.Getenv("CONFIG"))
	if errConfig != nil {
		panic(errConfig)
	}
	handler := factory.NewCommandHandlerBuilder().
		WithEventStore(eventStore).
		WithEventsEmitter(&pkg.StandardEventsEmitter{}).
		Build()
	err := handler.Execute(command, factory.NewShapeCommandApplier(eventStore))
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
