package grpc

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"trackclear.be/example/shape/domain/adapter"
)

type InMemoryGrpcEventStoreTestSuite struct {
	suite.Suite
	eventstore adapter.EventStore
	event      StandardEvent
}

func TestInMemoryGrpcEventStoreTestSuite(t *testing.T) {
	t.Skip()
	testingSuite := new(InMemoryGrpcEventStoreTestSuite)
	testingSuite.eventstore = NewInMemoryGrpcEventStore()
	testingSuite.event = newStandardEvent()
	suite.Run(t, testingSuite)
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
