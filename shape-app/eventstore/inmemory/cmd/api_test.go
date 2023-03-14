package cmd

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/eventstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InMemoryEventStoreTestSuite struct {
	suite.Suite
	eventstore eventstore.EventStore
}

func TestInMemoryEventStoreTestSuite(t *testing.T) {
	testingSuite := new(InMemoryEventStoreTestSuite)
	testingSuite.eventstore = NewInMemoryEventStore()
	suite.Run(t, testingSuite)
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_Save() {
	test := newEventForTest()
	_ = suite.eventstore.Save(test)
	events, _ := suite.eventstore.Load(test.AggregateId())
	assert.Contains(suite.T(), events, test)
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_LoadById() {
	test := newEventForTest()
	_ = suite.eventstore.Save(test)
	events, _ := suite.eventstore.Load(uuid.New())
	assert.Empty(suite.T(), events)
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_ErrorWhenNotFound() {
	_, err := suite.eventstore.Load(uuid.New())
	assert.Error(suite.T(), err)
}

type EventForTest struct {
	Id uuid.UUID
}

func newEventForTest() EventForTest {
	event := EventForTest{}
	event.Id = uuid.New()
	return event
}

func (e EventForTest) AggregateId() uuid.UUID {
	return e.Id
}

func (e EventForTest) Name() string {
	return "TEST_EVENT"
}

func (e EventForTest) Data() []byte {
	marshal, _ := json.Marshal(e)
	return marshal
}
