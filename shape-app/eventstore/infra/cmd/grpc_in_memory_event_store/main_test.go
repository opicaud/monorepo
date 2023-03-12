package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/eventstore"
	grpcEventStore "github.com/opicaud/monorepo/shape-app/eventstore/infra/pkg/grpc_in_memory_event_store"
	pb "github.com/opicaud/monorepo/shape-app/eventstore/infra/pkg/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

type InMemoryGrpcEventStoreTestSuite struct {
	suite.Suite
	eventstore eventstore.EventStore
	event      grpcEventStore.StandardEvent
}

func TestInMemoryGrpcEventStoreTestSuite(t *testing.T) {
	testingSuite := new(InMemoryGrpcEventStoreTestSuite)
	testingSuite.eventstore = grpcEventStore.NewInMemoryGrpcEventStore()
	testingSuite.event = grpcEventStore.NewStandardEvent("TEST")
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
