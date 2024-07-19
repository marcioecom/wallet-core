package createtransaction

import (
	"context"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/marcioecom/wallet-core/internal/gateway"
	"github.com/marcioecom/wallet-core/pkg/events"
	"github.com/marcioecom/wallet-core/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountFromID string  `json:"accountFromId"`
	AccountToID   string  `json:"accountToId"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"accountFromId"`
	AccountIDTo   string  `json:"accountToId"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcher
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcher,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (c *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	var output *CreateTransactionOutputDTO

	err := c.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := c.getAccountRepository(ctx)
		transactionRepository := c.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountFromID)
		if err != nil {
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountToID)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		if err = accountRepository.UpdateBalance(accountFrom); err != nil {
			return err
		}

		if err = accountRepository.UpdateBalance(accountTo); err != nil {
			return err
		}

		if err = transactionRepository.Create(transaction); err != nil {
			return err
		}
		output = &CreateTransactionOutputDTO{
			ID:            transaction.ID,
			AccountIDFrom: transaction.AccountFrom.ID,
			AccountIDTo:   transaction.AccountTo.ID,
			Amount:        transaction.Amount,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	c.TransactionCreated.SetPayload(output)
	c.EventDispatcher.Dispatch(c.TransactionCreated)

	return output, nil
}

func (c *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := c.Uow.GetRepository(ctx, "account")
	if err != nil {
		panic(err)
	}

	return repo.(gateway.AccountGateway)
}

func (c *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := c.Uow.GetRepository(ctx, "transaction")
	if err != nil {
		panic(err)
	}

	return repo.(gateway.TransactionGateway)
}
