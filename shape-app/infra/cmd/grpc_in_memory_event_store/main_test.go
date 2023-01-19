package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"trackclear.be/example/shape/domain/adapter"
	grpcEventStore "trackclear.be/example/shape/infra/pkg/inmemory_eventstore"
	pb "trackclear.be/example/shape/infra/pkg/proto"
)

type InMemoryGrpcEventStoreTestSuite struct {
	suite.Suite
	eventstore adapter.EventStore
	event      grpcEventStore.StandardEvent
}

func TestInMemoryGrpcEventStoreTestSuite(t *testing.T) {
	testingSuite := new(InMemoryGrpcEventStoreTestSuite)
	testingSuite.eventstore = grpcEventStore.NewInMemoryGrpcEventStore()
	testingSuite.event = newStandardEvent()
	go start()
	suite.Run(t, testingSuite)
}

func start() {
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

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryGrpcEventStoreSave() {
	err := suite.eventstore.Save(suite.event)
	assert.NoError(suite.T(), err)
}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryeGrpcEventStoreLoadKnownId() {
	events, err := suite.eventstore.Load(suite.event.AggregateId())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), events, 1)

}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryEventstoreErrorWhenUnknownId() {
	_, err := suite.eventstore.Load(uuid.New())
	assert.Error(suite.T(), err)
}

func newStandardEvent() grpcEventStore.StandardEvent {
	event := grpcEventStore.StandardEvent{}
	event.Id = uuid.New()
	return event
}
