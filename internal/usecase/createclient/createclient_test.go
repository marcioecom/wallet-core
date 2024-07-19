package createclient

import (
	"testing"

	"github.com/marcioecom/wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil).Times(1)

	input := CreateClientInputDTO{
		Name:  "John Doe",
		Email: "johndoe@mail.com",
	}

	usecase := NewCreateClientUseCase(m)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, output.Name, input.Name)
	assert.Equal(t, output.Email, input.Email)
	m.AssertExpectations(t)
}
