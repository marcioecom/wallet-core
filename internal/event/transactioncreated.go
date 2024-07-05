package event

import "time"

type TransactionCreated[T any] struct {
	Name    string
	Payload T
}

func NewTransactionCreated[T any]() *TransactionCreated[T] {
	return &TransactionCreated[T]{
		Name: "TransactionCreated",
	}
}

func (t *TransactionCreated[T]) GetName() string {
	return t.Name
}

func (t *TransactionCreated[T]) GetDateTime() time.Time {
	return time.Now()
}

func (t *TransactionCreated[T]) GetPayload() T {
	return t.Payload
}

func (t *TransactionCreated[T]) SetPayload(payload T) {
	t.Payload = payload
}
