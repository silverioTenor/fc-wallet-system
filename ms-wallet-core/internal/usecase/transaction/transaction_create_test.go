package transaction

import (
	"testing"

	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	fake "github.com/silverioTenor/fc-wallet-system/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientOne, _ := entity.NewClient("Willy Wonka", "willy@wonka.com")
	accountOne, _ := entity.NewAccount(clientOne)
	accountOne.Credit(1000)

	clientTwo, _ := entity.NewClient("Jane Doe", "j@doe.com")
	accountTwo, _ := entity.NewAccount(clientTwo)
	accountTwo.Credit(1000)

	mockAccountGateway := &fake.AccountGatewayMock{}
	mockAccountGateway.On("FindByID", accountOne.ID).Return(accountOne, nil)
	mockAccountGateway.On("FindByID", accountTwo.ID).Return(accountTwo, nil)

	mockTransactionGateway := &fake.TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	inputDTO := InputCreateTransactionDTO{
		AccountIDFrom: accountOne.ID,
		AccountIDTo: accountTwo.ID,
		Amount: 100,
	}

	usecase := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway)
	outputDTO, err := usecase.Execute(inputDTO)

	assert.Nil(t, err)
	assert.NotNil(t, outputDTO)
	mockAccountGateway.AssertExpectations(t)
	mockTransactionGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)
}