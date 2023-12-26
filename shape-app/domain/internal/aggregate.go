package internal

import (
	"github.com/opicaud/monorepo/events/pkg"
)

type Shape interface {
	HandleCreationCommand(command CreationCommand) Created
	HandleStretchCommand(command StretchCommand) Stretched
	ApplyCreatedEvent(area Created) Shape
	ApplyStretchedEvent(area Stretched) Shape
}

type CommandApplier interface {
	ApplyCreationCommand(command CreationCommand) ([]pkg.DomainEvent, error)
	ApplyStretchCommand(command StretchCommand) ([]pkg.DomainEvent, error)
}
