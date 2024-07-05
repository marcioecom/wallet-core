package createtransaction

import (
	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/marcioecom/wallet-core/internal/gateway"
	"github.com/marcioecom/wallet-core/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
	EventDispatcher    events.EventDispatcher
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcher,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := c.AccountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.AccountGateway.FindByID(input.AccountToID)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	if err = c.TransactionGateway.Create(transaction); err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	c.TransactionCreated.SetPayload(output)
	c.EventDispatcher.Dispatch(c.TransactionCreated)

	return output, nil
}
