package events

import "time"

type EventInterface[T any] interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() T
	SetPayload(payload T)
}

type EventHandlerInterface[T any] interface {
	Handle(event EventInterface[T])
}

type EventDispatcher[T any] interface {
	Register(eventName string, handler EventHandlerInterface[T]) error
	Dispatch(event EventInterface[T]) error
	Remove(eventName string, handler EventHandlerInterface[T]) error
	Has(eventName string, handler EventHandlerInterface[T]) bool
	Clear()
}
