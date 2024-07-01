package createtransaction

import (
	"testing"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

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

	am := &AccountGatewayMock{}
	am.On("FindByID", account1.ID).Return(account1, nil).Times(1)
	am.On("FindByID", account2.ID).Return(account2, nil).Times(1)

	tm := &TransactionGatewayMock{}
	tm.On("Create", mock.Anything).Return(nil).Times(1)

	input := CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100,
	}

	usecase := NewCreateTransactionUseCase(tm, am)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	assert.Equal(t, account1.Balance, float64(0))
	assert.Equal(t, account2.Balance, float64(100))

	tm.AssertExpectations(t)
	am.AssertExpectations(t)
}
