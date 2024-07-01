package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validAccounts() (*Account, *Account) {
	client1, _ := NewClient("Client 1", "client1@mail.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Client 2", "client2@mail.com")
	account2, _ := NewAccount(client2)
	return account1, account2
}

func TestCreateTransaction(t *testing.T) {
	account1, account2 := validAccounts()

	account1.Credit(1000)
	account2.Credit(1000)

	tx, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, account1.Balance, float64(900))
	assert.Equal(t, account2.Balance, float64(1100))
}

func TestCreateTransactionWithAmountZero(t *testing.T) {
	account1, account2 := validAccounts()

	account1.Credit(100)
	account2.Credit(100)

	tx, err := NewTransaction(account1, account2, 0)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Nil(t, tx)
	assert.Equal(t, account1.Balance, float64(100))
	assert.Equal(t, account2.Balance, float64(100))
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	account1, account2 := validAccounts()

	account1.Credit(100)
	account2.Credit(100)

	tx, err := NewTransaction(account1, account2, 150)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, tx)
	assert.Equal(t, account1.Balance, float64(100))
	assert.Equal(t, account2.Balance, float64(100))
}
