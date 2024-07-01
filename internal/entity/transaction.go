package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string `json:"id"`
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	tx := &Transaction{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	if err := tx.Validate(); err != nil {
		return nil, err
	}

	tx.Commit()

	return tx, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	return nil
}
