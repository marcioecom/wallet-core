package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload any
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (t *TransactionCreated) GetName() string {
	return t.Name
}

func (t *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}

func (t *TransactionCreated) GetPayload() any {
	return t.Payload
}

func (t *TransactionCreated) SetPayload(payload any) {
	t.Payload = payload
}
