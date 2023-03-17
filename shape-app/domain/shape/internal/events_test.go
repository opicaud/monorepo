package internal

import (
	"github.com/google/uuid"
	"github.com/smartystreets/assertions"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactoryEventTestSuite struct {
	suite.Suite
}

func TestFactoryEventTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryEventTestSuite))
}

func (suite *FactoryEventTestSuite) TestCreateEvent() {
	factory := NewShapeEventFactory()
	u := uuid.New()
	event := factory.NewShapeCreatedEvent(u, "nature", 1, 2, 3)
	domainEvent := factory.NewDeserializedEvent(u, event).(*Created)

	assertions.ShouldEqual(domainEvent.AggregateId(), u)
	assertions.ShouldEqual(domainEvent.Name(), event.Name())
	assertions.ShouldEqual(domainEvent.Id, u)
	assertions.ShouldEqual(domainEvent.Nature, "nature")
	assertions.ShouldEqual(domainEvent.Area, 1)
	assertions.ShouldEqual(domainEvent.Dimensions[0], 2)
	assertions.ShouldEqual(domainEvent.Dimensions[1], 3)

}

func (suite *FactoryEventTestSuite) TestStretchedEvent() {
	factory := NewShapeEventFactory()
	u := uuid.New()
	event := factory.NewShapeStretchedEvent(u, 1, 2, 3)
	domainEvent := factory.NewDeserializedEvent(u, event).(*Stretched)

	assertions.ShouldEqual(domainEvent.AggregateId(), u)
	assertions.ShouldEqual(domainEvent.Name(), event.Name())
	assertions.ShouldEqual(domainEvent.Id, u)
	assertions.ShouldEqual(domainEvent.Area, 1)
	assertions.ShouldEqual(domainEvent.Dimensions[0], 2)
	assertions.ShouldEqual(domainEvent.Dimensions[1], 3)

}
