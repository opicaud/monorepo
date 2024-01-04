package internal

import (
	cqrs "github.com/opicaud/monorepo/cqrs/v3/pkg"
)

type Shape interface {
	HandleCreationCommand(command CreationCommand) Created
	HandleStretchCommand(command StretchCommand) Stretched
	ApplyCreatedEvent(area Created) Shape
	ApplyStretchedEvent(area Stretched) Shape
}

type CommandApplier interface {
	ApplyCreationCommand(command CreationCommand) ([]cqrs.DomainEvent, error)
	ApplyStretchCommand(command StretchCommand) ([]cqrs.DomainEvent, error)
}
