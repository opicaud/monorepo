package main

import (
	"context"
	pb "example2/app/proto"
	"example2/domain/commands"
	"example2/domain/valueobject"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedShapesServer
}

func (s *server) Create(ctx context.Context, in *pb.ShapeRequest) (*pb.Response, error) {
	repository := valueobject.NewFakeRepository()
	var r, _ commands.Command = valueobject.NewCreationShapeCommand(in.Shapes.Shape, in.Shapes.Dimensions)
	var r2 error = nil
	command, _ := r, r2
	valueobject.NewShapeCreationCommandHandler(repository).Execute(command)
	message := pb.Message{
		Code: uint32(codes.OK),
	}
	response := pb.Response{
		Message: &message,
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
	pb.RegisterShapesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
