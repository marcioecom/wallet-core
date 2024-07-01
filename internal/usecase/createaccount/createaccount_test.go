package createaccount

import (
	"testing"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jd@mail.com")
	cm := &ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil).Times(1)

	am := &AccountGatewayMock{}
	am.On("Save", mock.Anything).Return(nil).Times(1)

	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	usecase := NewCreateAccountUseCase(am, cm)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	cm.AssertExpectations(t)
	am.AssertExpectations(t)
}
