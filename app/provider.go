package main

import (
	"context"
	pb "example2/app/proto"
	"example2/domain/commands/factory"
	"example2/domain/commands/shapecreationcommand"
	"example2/domain/utils"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) CalculateArea(ctx context.Context, in *pb.AreaRequest) (*pb.AreaResponse, error) {
	repository := utils.NewFakeRepository()
	command, _ := factory.NewFactory().NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions...)
	shapecreationcommand.NewShapeCreationCommandHandler(repository).Execute(command)
	response := pb.AreaResponse{
		Value: repository.Get(0).GetArea(),
	}
	return &response, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
