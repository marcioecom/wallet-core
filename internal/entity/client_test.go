package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "jd@mail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.Name, "John Doe")
	assert.Equal(t, client.Email, "jd@mail.com")
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.Error(t, err, "name is required")
	assert.Nil(t, client)

	client, err = NewClient("John Doe", "")
	assert.Error(t, err, "email is required")
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")
	err := client.Update("John Doe Update", "jdupdate@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, client.Name, "John Doe Update")
	assert.Equal(t, client.Email, "jdupdate@mail.com")
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")
	err := client.Update("", "jdupdate@mail.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@mail.com")
	account, _ := NewAccount(client)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, len(client.Accounts), 1)
	assert.Equal(t, client.Accounts[0], account)
}

func TestAddAccountThatNotBelongsToClient(t *testing.T) {
	client1, _ := NewClient("John Doe 1", "jd1@mail.com")
	client2, _ := NewClient("John Doe 2", "jd2@mail.com")
	account, _ := NewAccount(client1)

	err := client2.AddAccount(account)
	assert.Error(t, err, "account does not belong to the client")
}
