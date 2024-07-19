package createaccount

import (
	"testing"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/marcioecom/wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jd@mail.com")
	cm := &mocks.ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil).Times(1)

	am := &mocks.AccountGatewayMock{}
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
