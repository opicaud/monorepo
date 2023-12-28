package internal

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/smarty/assertions"
	"github.com/stretchr/testify/suite"
	"log"
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
	u, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	event := factory.NewShapeCreatedEvent(u, "square", 1, 2, 3)
	domainEvent := factory.NewDeserializedEvent(u, event).(*Created)
	log.Println(base64.StdEncoding.EncodeToString(event.Data()))
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
