package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")

	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, account.Balance, float64(0))
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.Error(t, err, "client is nil")
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")
	account, _ := NewAccount(client)

	account.Credit(100)
	assert.Equal(t, account.Balance, float64(100))
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")
	account, _ := NewAccount(client)

	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, account.Balance, float64(50))
}
