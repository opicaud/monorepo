package aggregate

import (
	"example2/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (suite *CommandHandlerTestSuite) TestHandlerAShapeCreationCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	err := suite.handler.Execute(newCreationShapeCommand(nature, dimensions))

	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	assert.Equal(suite.T(), ShapeCreated{id: suite.subscriber.ids[0], Nature: nature, dimensions: dimensions, Area: 2}, suite.subscriber.events[0])
	assert.NoError(suite.T(), err)

}

func (suite *CommandHandlerTestSuite) TestHandlerAStretchCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	assert.NoError(suite.T(), suite.handler.Execute(newCreationShapeCommand(nature, dimensions)))
	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	id := suite.subscriber.ids[0]
	err := suite.handler.Execute(newStrechShapeCommand(id, 2))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.subscriber.events))

	assert.Equal(suite.T(), ShapeStreched{id: id, Area: 8, dimensions: []float32{2, 4}}, suite.subscriber.events[1])

}

type CommandHandlerTestSuite struct {
	suite.Suite
	handler    ShapeCommandHandler
	subscriber SubscriberForTest
	infra      infra.Provider
}

// this function executes before each test case
func (suite *CommandHandlerTestSuite) SetupTest() {
	suite.subscriber = SubscriberForTest{}
	suite.infra = infra.NewInfraBuilder().
		WithEventStore(infra.NewInMemoryEventStore()).
		WithEmitter(&infra.StandardEventsEmitter{}).
		Build()
	suite.handler = NewShapeCreationCommandHandlerBuilder().WithInfraProvider(suite.infra).WithSubscriber(&suite.subscriber).Build()
}

func TestRunCommandHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CommandHandlerTestSuite))
}

type SubscriberForTest struct {
	events []infra.Event
	ids    []uuid.UUID
}

func (s *SubscriberForTest) Update(events []infra.Event) {
	s.events = append(s.events, events...)
	s.ids = []uuid.UUID{}
	for _, e := range s.events {
		s.ids = append(s.ids, e.AggregateId())
	}
}
