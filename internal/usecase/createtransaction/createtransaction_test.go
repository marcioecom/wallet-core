package createtransaction

import (
	"context"
	"testing"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/marcioecom/wallet-core/internal/event"
	"github.com/marcioecom/wallet-core/internal/usecase/mocks"
	"github.com/marcioecom/wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func validAccounts() (*entity.Account, *entity.Account) {
	client1, _ := entity.NewClient("Client 1", "client1@mail.com")
	account1, _ := entity.NewAccount(client1)
	client2, _ := entity.NewClient("Client 2", "client2@mail.com")
	account2, _ := entity.NewAccount(client2)
	return account1, account2
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	account1, account2 := validAccounts()
	account1.Credit(100)

	am := &mocks.AccountGatewayMock{}
	am.On("FindByID", account1.ID).Return(account1, nil).Times(1)
	am.On("FindByID", account2.ID).Return(account2, nil).Times(1)

	tm := &mocks.TransactionGatewayMock{}

	input := CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100,
	}

	event := event.NewTransactionCreated()
	dispatcher := events.NewEventDispatcher()

	ctx := context.Background()
	uow := &mocks.UowMock{}
	uow.On("Do", ctx, mock.Anything).Return(nil).Times(1)
	uow.On("GetRepository", ctx, "account").Return(am, nil).Times(1)
	uow.On("GetRepository", ctx, "transaction").Return(tm, nil).Times(1)

	usecase := NewCreateTransactionUseCase(uow, dispatcher, event)
	output, err := usecase.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	assert.Equal(t, account1.Balance, float64(0))
	assert.Equal(t, account2.Balance, float64(100))

	uow.AssertExpectations(t)
}
