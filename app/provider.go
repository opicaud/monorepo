package main

import (
	"context"
	pb "example2/app/proto"
	"example2/domain/commands"
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

func (s *server) CalculateMultiV2(ctx context.Context, in *pb.AreaRequestV2) (*pb.AreaResponse, error) {
	handler := commands.NewFullShapeCommandHandler(utils.NewFakeRepository())
	factory := commands.NewFactory()
	command, _ := factory.CreateAFullShapeCommand("nature", 1, 2)
	handler.Execute(command)
	return nil, nil
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
