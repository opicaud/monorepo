package shape

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/opicaud/monorepo/shape-app/domain/adapter"
)

func (suite *CommandHandlerTestSuite) TestHandlerAShapeCreationCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	err := suite.handler.Execute(newCreationShapeCommand(nature, dimensions))

	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	assert.Equal(suite.T(), Created{id: suite.subscriber.ids[0], Nature: nature, dimensions: dimensions, Area: 2}, suite.subscriber.events[0])
	assert.NoError(suite.T(), err)

}

func (suite *CommandHandlerTestSuite) TestHandlerAStretchCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	assert.NoError(suite.T(), suite.handler.Execute(newCreationShapeCommand(nature, dimensions)))
	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	id := suite.subscriber.ids[0]
	err := suite.handler.Execute(newStretchShapeCommand(id, 2))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.subscriber.events))

	assert.Equal(suite.T(), Stretched{id: id, Area: 8, dimensions: []float32{2, 4}}, suite.subscriber.events[1])

}

func (suite *CommandHandlerTestSuite) TestHandleStretchWithAreaNotFound() {

	err := suite.handler.Execute(newStretchShapeCommand(uuid.New(), 2))
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(suite.subscriber.events))

}

type CommandHandlerTestSuite struct {
	suite.Suite
	handler    CommandHandler
	subscriber SubscriberForTest
	infra      adapter.Provider
}

// this function executes before each test case
func (suite *CommandHandlerTestSuite) SetupTest() {
	suite.subscriber = SubscriberForTest{}
	suite.infra = adapter.NewInfraBuilder().
		WithEventStore(adapter.NewInMemoryEventStore()).Build()
	suite.handler = NewShapeCreationCommandHandlerBuilder().WithInfraProvider(suite.infra).WithSubscriber(&suite.subscriber).Build()
}

func TestRunCommandHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CommandHandlerTestSuite))
}

type SubscriberForTest struct {
	events []adapter.DomainEvent
	ids    []uuid.UUID
}

func (s *SubscriberForTest) Update(events []adapter.DomainEvent) {
	s.events = append(s.events, events...)
	s.ids = []uuid.UUID{}
	for _, e := range s.events {
		s.ids = append(s.ids, e.AggregateId())
	}
}
