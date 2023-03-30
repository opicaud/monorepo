package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedEventStoreServer
	events []*pb.Event
}

func (s *server) Save(ctx context.Context, in *pb.Events) (*pb.Response, error) {
	log.Printf("received %s\n", in)
	s.events = append(s.events, in.GetEvent()...)
	return &pb.Response{}, nil
}
func (s *server) Load(ctx context.Context, in *pb.UUID) (*pb.Response, error) {
	log.Printf("load %s\n", in)
	events := s.search(in.Id)
	r := &pb.Response{Events: &pb.Events{Event: events}}
	return r, nil
}

func (s *server) search(id string) []*pb.Event {
	w := 0
	for _, e := range s.events {
		if e.AggregateId.Id == id {
			s.events[w] = e
			w++
		}
	}
	return s.events[0:w]

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
	pb.RegisterEventStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
