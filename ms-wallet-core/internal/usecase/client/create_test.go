package client

import (
	"testing"

	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	usecase := NewCreateClientUseCase(clientGatewayMock)

	input := CreateClientInputDTO{
		Name: "Willy Wonka",
		Email: "willy@wonka.com",
	}
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Willy Wonka", output.Name)
	assert.Equal(t, "willy@wonka.com", output.Email)

	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}