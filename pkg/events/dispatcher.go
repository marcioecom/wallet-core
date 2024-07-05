package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type Dispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (d *Dispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if d.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

func (d *Dispatcher) Dispatch(event EventInterface) error {
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

func (d *Dispatcher) Remove(eventName string, handler EventHandlerInterface) error {
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

func (d *Dispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if handlers, ok := d.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (d *Dispatcher) Clear() {
	d.handlers = make(map[string][]EventHandlerInterface)
}
