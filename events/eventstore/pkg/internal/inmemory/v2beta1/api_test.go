package inmemory

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg/v2beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/trace"
	"testing"
)

type InMemoryEventStoreTestSuite struct {
	suite.Suite
	eventStore pkg.EventStore
}

func TestInMemoryEventStoreTestSuite(t *testing.T) {
	testingSuite := new(InMemoryEventStoreTestSuite)
	testingSuite.eventStore = NewInMemoryEventStore()
	suite.Run(t, testingSuite)
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_Save() {
	test := newEventForTest()
	ctx, domainEvents, err := suite.eventStore.Save(context.Background(), test)
	assert.Len(suite.T(), domainEvents, 1)
	assert.Equal(suite.T(), domainEvents[0], test)
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), trace.SpanContextFromContext(ctx).IsValid())
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_LoadById() {
	test := newEventForTest()
	background := context.Background()
	_, _, _ = suite.eventStore.Save(background, test)
	context, events, _ := suite.eventStore.Load(background, uuid.New())
	assert.True(suite.T(), trace.SpanContextFromContext(context).IsValid())
	assert.Empty(suite.T(), events)
}

func (suite *InMemoryEventStoreTestSuite) TestInMemoryEventStore_ErrorWhenNotFound() {
	background := context.Background()
	ctx, _, err := suite.eventStore.Load(background, uuid.New())
	assert.True(suite.T(), trace.SpanFromContext(ctx).SpanContext().IsValid())
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
