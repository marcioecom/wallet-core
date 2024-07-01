package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string `json:"id"`
	Client    *Client
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(client *Client) (*Account, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}

	a := &Account{
		ID:        uuid.NewString(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return a, nil
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}
