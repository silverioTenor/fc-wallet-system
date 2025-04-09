package client

import (
	"testing"

	fake "github.com/silverioTenor/fc-wallet-system/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGatewayMock := &fake.ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	usecase := NewCreateClientUseCase(clientGatewayMock)

	input := InputCreateClientDTO{
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