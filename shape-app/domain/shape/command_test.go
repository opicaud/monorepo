package shape

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/eventstore/inmemory/cmd"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/cqrs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestFullShapeCommand(t *testing.T) {
	command := newCreationShapeCommand("rectangle", []float32{1, 3})
	assert.NotNil(t, command)
}

func TestStretchShapeCommand(t *testing.T) {
	id := uuid.New()
	command := newStretchShapeCommand(id, 1)
	assert.Equal(t, id, command.id)
	assert.Equal(t, float32(1), command.stretchBy)
}

func (suite *CommandHandlerTestSuite) TestHandlerAShapeCreationCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	err := suite.handler.Execute(newCreationShapeCommand(nature, dimensions), NewShapeCommandApplier(suite.eventsFramework))

	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	assert.Equal(suite.T(), Created{id: suite.subscriber.ids[0], Nature: nature, Dimensions: dimensions, Area: 2}, suite.subscriber.events[0])
	assert.NoError(suite.T(), err)

}

func (suite *CommandHandlerTestSuite) TestHandlerAStretchCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	applier := NewShapeCommandApplier(suite.eventsFramework)
	assert.NoError(suite.T(), suite.handler.Execute(newCreationShapeCommand(nature, dimensions), applier))
	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	id := suite.subscriber.ids[0]
	err := suite.handler.Execute(newStretchShapeCommand(id, 2), applier)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.subscriber.events))

	assert.Equal(suite.T(), Stretched{id: id, Area: 8, Dimensions: []float32{2, 4}}, suite.subscriber.events[1])

}

func (suite *CommandHandlerTestSuite) TestHandleStretchWithAreaNotFound() {
	applier := NewShapeCommandApplier(suite.eventsFramework)
	err := suite.handler.Execute(newStretchShapeCommand(uuid.New(), 2), applier)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(suite.subscriber.events))

}

type CommandHandlerTestSuite struct {
	suite.Suite
	handler         cqrs.CommandHandler[cqrs.Command[CommandApplier], CommandApplier]
	subscriber      SubscriberForTest
	eventsFramework pkg.Provider
}

// this function executes before each test case
func (suite *CommandHandlerTestSuite) SetupTest() {
	suite.subscriber = SubscriberForTest{}
	suite.eventsFramework = pkg.NewEventsFrameworkBuilder().
		WithEventStore(cmd.NewInMemoryEventStore()).Build()
	suite.handler = NewCommandHandlerBuilder().
		WithEventsFramework(suite.eventsFramework).
		WithSubscriber(&suite.subscriber).Build()
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
