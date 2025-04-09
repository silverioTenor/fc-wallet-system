package account

import (
	"testing"

	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	fake "github.com/silverioTenor/fc-wallet-system/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Willy Wonka", "willy@wonka.com")
	clientMock := &fake.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &fake.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := InputCreateAccountDTO{
		ClientID: client.ID,
	}
	output, err := usecase.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
