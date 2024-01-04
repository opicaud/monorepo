package internal

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
	"log"
)

type ShapeEventFactory struct{}

func NewShapeEventFactory() ShapeEventFactory {
	return ShapeEventFactory{}
}

func (f ShapeEventFactory) NewShapeCreatedEvent(id uuid.UUID, nature string, area float32, dimensions ...float32) Created {
	return Created{Id: id, Dimensions: dimensions, Area: area, Nature: nature}
}

func (f ShapeEventFactory) NewShapeStretchedEvent(id uuid.UUID, area float32, dimensions ...float32) Stretched {
	return Stretched{Id: id, Dimensions: dimensions, Area: area}
}

func (f ShapeEventFactory) NewDeserializedEvent(aggregateId uuid.UUID, event cqrs.DomainEvent) cqrs.DomainEvent {
	switch event.Name() {
	case "SHAPE_CREATED":
		v := &Created{}
		err := json.Unmarshal(event.Data(), v)
		if err != nil {
			log.Panic(err)
		}
		v.Id = aggregateId
		return v
	case "SHAPE_STRETCHED":
		v := &Stretched{}
		_ = json.Unmarshal(event.Data(), v)
		v.Id = aggregateId
		return v
	default:
		panic(fmt.Errorf("%s is not known as event", event.Name()))
	}

}

type Stretched struct {
	Area       float32
	Dimensions []float32
	Id         uuid.UUID
}

func (a Stretched) Data() []byte {
	marshal, _ := json.Marshal(a)
	return marshal
}

type Created struct {
	Nature     string
	Dimensions []float32
	Id         uuid.UUID
	Area       float32
}

func (s Created) AggregateId() uuid.UUID {
	return s.Id
}

func (s Created) Name() string {
	return "SHAPE_CREATED"
}

func (s Created) Data() []byte {
	marshal, _ := json.Marshal(s)
	return marshal
}

func (s Created) ApplyOn(shape Shape) Shape {
	return shape.ApplyCreatedEvent(s)
}

func (a Stretched) AggregateId() uuid.UUID {
	return a.Id
}

func (a Stretched) Name() string {
	return "SHAPE_STRETCHED"
}

func (a Stretched) ApplyOn(shape Shape) Shape {
	return shape.ApplyStretchedEvent(a)
}
