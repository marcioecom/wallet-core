package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type Dispatcher[T any] struct {
	handlers map[string][]EventHandlerInterface[T]
}

func NewEventDispatcher[T any]() *Dispatcher[T] {
	return &Dispatcher[T]{
		handlers: make(map[string][]EventHandlerInterface[T]),
	}
}

func (d *Dispatcher[T]) Register(eventName string, handler EventHandlerInterface[T]) error {
	if d.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

func (d *Dispatcher[T]) Dispatch(event EventInterface[T]) error {
	var wg sync.WaitGroup

	if handlers, ok := d.handlers[event.GetName()]; ok {
		for _, h := range handlers {
			wg.Add(1)
			go func() {
				h.Handle(event)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	return nil
}

func (d *Dispatcher[T]) Remove(eventName string, handler EventHandlerInterface[T]) error {
	if handlers, ok := d.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				d.handlers[eventName] = append(d.handlers[eventName][:i], d.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (d *Dispatcher[T]) Has(eventName string, handler EventHandlerInterface[T]) bool {
	if handlers, ok := d.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (d *Dispatcher[T]) Clear() {
	d.handlers = make(map[string][]EventHandlerInterface[T])
}
