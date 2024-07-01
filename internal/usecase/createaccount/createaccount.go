package createaccount

import (
	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/marcioecom/wallet-core/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (c *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := c.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	if err = c.AccountGateway.Save(account); err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
