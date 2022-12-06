package infra

type EventStore interface {
	Save(events ...Event) error
}
