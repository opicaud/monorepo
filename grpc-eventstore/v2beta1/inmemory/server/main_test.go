package main

import (
	"context"
	"github.com/google/uuid"
	pkg "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
	pkg2 "github.com/opicaud/monorepo/grpc-eventstore/v2beta1/inmemory/client"
	internal "github.com/opicaud/monorepo/grpc-eventstore/v2beta1/inmemory/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"testing"
)

type InMemoryGrpcEventStoreTestSuite struct {
	suite.Suite
	eventstore pkg.EventStore
	event      internal.StandardEvent
}

func TestInMemoryGrpcEventStoreTestSuite(t *testing.T) {
	result := make(chan int, 1)
	go start(result)
	port := <-result
	close(result)
	testingSuite := new(InMemoryGrpcEventStoreTestSuite)
	testingSuite.eventstore = pkg2.NewInMemoryGrpcEventStoreFrom("localhost", port)
	testingSuite.event = internal.NewStandardEventForTest("TEST")
	_, client := testingSuite.eventstore.GetHealthClient(context.Background())
	check, err := client.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
	assert.NoError(t, err)
	assert.Equal(t, check.Status, grpc_health_v1.HealthCheckResponse_SERVING)
	suite.Run(t, testingSuite)
}

func start(result chan int) {
	lis, err := net.Listen("tcp", ":0")
	s := startServer(err)
	result <- lis.Addr().(*net.TCPAddr).Port
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryGrpcEventStoreSave() {
	_, _, err := suite.eventstore.Save(context.Background(), suite.event)
	assert.NoError(suite.T(), err)
}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryeGrpcEventStoreLoadKnownId() {
	_, events, err := suite.eventstore.Load(context.Background(), suite.event.AggregateId())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), events, 1)

}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryEventstoreErrorWhenUnknownId() {
	_, events, err := suite.eventstore.Load(context.Background(), uuid.New())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), events, 0)
}

func (suite *InMemoryGrpcEventStoreTestSuite) TestInMemoryEventstoreRemoveEvent() {
	ctx, _, _ := suite.eventstore.Save(context.Background(), suite.event)
	ctx, err := suite.eventstore.Remove(ctx, suite.event.AggregateId())
	assert.NoError(suite.T(), err)
	ctx, events, _ := suite.eventstore.Load(ctx, suite.event.AggregateId())
	assert.Len(suite.T(), events, 0)
}
