package pkg

import (
	"context"
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/v3/pkg"
	pkg "github.com/opicaud/monorepo/grpc-eventstore/v2/cmd"
	"github.com/opicaud/monorepo/shape-app/domain/internal"
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
	_, err := suite.handler.Execute(context.TODO(), internal.NewCreationShapeCommand(nature, dimensions), internal.NewShapeCommandApplier(suite.eventStore))

	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	assert.Equal(suite.T(), internal.Created{Id: suite.subscriber.ids[0], Nature: nature, Dimensions: dimensions, Area: 2}, suite.subscriber.events[0])
	assert.NoError(suite.T(), err)

}

func (suite *CommandHandlerTestSuite) TestHandlerAStretchCommand() {

	nature := "rectangle"
	dimensions := []float32{1, 2}
	applier := internal.NewShapeCommandApplier(suite.eventStore)
	_, err := suite.handler.Execute(context.TODO(), internal.NewCreationShapeCommand(nature, dimensions), applier)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(suite.subscriber.events))

	id := suite.subscriber.ids[0]
	_, err = suite.handler.Execute(context.TODO(), internal.NewStretchShapeCommand(id, 2), applier)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.subscriber.events))

	assert.Equal(suite.T(), internal.Stretched{Id: id, Area: 8, Dimensions: []float32{2, 4}}, suite.subscriber.events[1])

}

func (suite *CommandHandlerTestSuite) TestHandleStretchWithAreaNotFound() {
	applier := New().NewShapeCommandApplier(suite.eventStore)
	_, err := suite.handler.Execute(context.TODO(), New().NewStretchShapeCommand(uuid.New(), 2), applier)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, len(suite.subscriber.events))

}

type CommandHandlerTestSuite struct {
	suite.Suite
	handler    cqrs.CommandHandler[cqrs.Command[internal.CommandApplier], internal.CommandApplier]
	subscriber SubscriberForTest
	eventStore cqrs.EventStore
}

// this function executes before each test case
func (suite *CommandHandlerTestSuite) SetupTest() {
	suite.subscriber = SubscriberForTest{}
	suite.eventStore, _ = pkg.NewEventsFrameworkFromConfig("")
	suite.handler = New().NewCommandHandlerBuilder().
		WithEventStore(suite.eventStore).
		WithSubscriber(&suite.subscriber).
		WithEventsEmitter(&cqrs.StandardEventsEmitter{}).
		Build()
}

func TestRunCommandHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CommandHandlerTestSuite))
}

type SubscriberForTest struct {
	events []cqrs.DomainEvent
	ids    []uuid.UUID
}

func (s *SubscriberForTest) Update(ctx context.Context, eventsChn chan []cqrs.DomainEvent) context.Context {
	events := <-eventsChn
	s.events = append(s.events, events...)
	s.ids = []uuid.UUID{}
	for _, e := range s.events {
		s.ids = append(s.ids, e.AggregateId())
	}
	return ctx
}
