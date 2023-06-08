package main

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/internal"
	inmem "github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/pkg"
	pb "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	"github.com/opicaud/monorepo/events/pkg/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

type InMemoryGrpcEventStoreTestSuite struct {
	suite.Suite
	eventstore pkg.EventStore
	event      eventstore.StandardEvent
}

func TestInMemoryGrpcEventStoreTestSuite(t *testing.T) {
	result := make(chan int, 1)
	go start(result)
	port := <-result
	close(result)
	testingSuite := new(InMemoryGrpcEventStoreTestSuite)
	testingSuite.eventstore = inmem.NewInMemoryGrpcEventStoreFrom("localhost", port)
	testingSuite.event = eventstore.NewStandardEventForTest("TEST")
	suite.Run(t, testingSuite)
}

func start(result chan int) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEventStoreServer(s, &server{})
	result <- lis.Addr().(*net.TCPAddr).Port
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
	events, err := suite.eventstore.Load(uuid.New())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), events, 0)
}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryEventstoreRemoveEvent() {
	_ = suite.eventstore.Save(suite.event)
	err := suite.eventstore.Remove(suite.event.AggregateId())
	assert.NoError(suite.T(), err)
	events, _ := suite.eventstore.Load(suite.event.AggregateId())
	assert.Len(suite.T(), events, 0)
}
