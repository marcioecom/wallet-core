package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Accounts  []*Account
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewClient(name, email string) (*Client, error) {
	c := &Client{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (c *Client) Update(name, email string) error {
	c.Name = name
	c.Email = email

	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("account does not belong to the client")
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
