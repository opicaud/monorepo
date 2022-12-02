package infra

type EventBus interface {
	NotifyAll()
}
