package pkg

import (
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/pkg"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/domain/internal"
	"github.com/opicaud/monorepo/shape-app/domain/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactoryTestSuite struct {
	suite.Suite
}

func TestFactoryTestSuite2(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}

func (suite *FactoryTestSuite) TestCreateACommandFullShape() {
	var command = internal.NewCreationShapeCommand("a-shape", []float32{1, 2})
	assert.NotNil(suite.T(), command)
}

func (suite *CommandHandlerTestSuite) TestHandlerAShapeCreationCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	err := suite.handler.Execute(internal.NewCreationShapeCommand(nature, dimensions), internal.NewShapeCommandApplier(suite.eventStore))

	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	assert.Equal(suite.T(), internal.Created{Id: suite.subscriber.ids[0], Nature: nature, Dimensions: dimensions, Area: 2}, suite.subscriber.events[0])
	assert.NoError(suite.T(), err)

}

func (suite *CommandHandlerTestSuite) TestHandlerAStretchCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	applier := internal.NewShapeCommandApplier(suite.eventStore)
	assert.NoError(suite.T(), suite.handler.Execute(internal.NewCreationShapeCommand(nature, dimensions), applier))
	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	id := suite.subscriber.ids[0]
	err := suite.handler.Execute(internal.NewStretchShapeCommand(id, 2), applier)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.subscriber.events))

	assert.Equal(suite.T(), internal.Stretched{Id: id, Area: 8, Dimensions: []float32{2, 4}}, suite.subscriber.events[1])

}

func (suite *CommandHandlerTestSuite) TestHandleStretchWithAreaNotFound() {
	applier := New().NewShapeCommandApplier(suite.eventStore)
	err := suite.handler.Execute(New().NewStretchShapeCommand(uuid.New(), 2), applier)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(suite.subscriber.events))

}

type CommandHandlerTestSuite struct {
	suite.Suite
	handler    cqrs.CommandHandler[cqrs.Command[internal.CommandApplier], internal.CommandApplier]
	subscriber SubscriberForTest
	eventStore pkg.EventStore
}

// this function executes before each test case
func (suite *CommandHandlerTestSuite) SetupTest() {
	suite.subscriber = SubscriberForTest{}
	suite.eventStore = test.NewFakeInMemoryEventStore()
	suite.handler = New().NewCommandHandlerBuilder().
		WithEventStore(suite.eventStore).
		WithSubscriber(&suite.subscriber).
		WithEventsEmitter(&pkg.StandardEventsEmitter{}).
		Build()
}

func TestRunCommandHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CommandHandlerTestSuite))
}

type SubscriberForTest struct {
	events []pkg.DomainEvent
	ids    []uuid.UUID
}

func (s *SubscriberForTest) Update(events []pkg.DomainEvent) {
	s.events = append(s.events, events...)
	s.ids = []uuid.UUID{}
	for _, e := range s.events {
		s.ids = append(s.ids, e.AggregateId())
	}
}
