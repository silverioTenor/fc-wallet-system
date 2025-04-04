package account

import (
	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/silverioTenor/fc-wallet-system/internal/gateway"
)

type InputCreateAccountDTO struct {
	ClientID string
}

type OutputCreateAccountDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.IAccountGateway
	ClientGateway  gateway.IClientGateway
}

func NewCreateAccountUseCase(account gateway.IAccountGateway, client gateway.IClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: account,
		ClientGateway:  client,
	}
}

func (usecase *CreateAccountUseCase) Execute(input InputCreateAccountDTO) (*OutputCreateAccountDTO, error) {
	client, clientError := usecase.ClientGateway.Get(input.ClientID)
	if clientError != nil {
		return nil, clientError
	}

	account, accountError := entity.NewAccount(client)
	if accountError != nil {
		return nil, accountError
	}

	saveError := usecase.AccountGateway.Save(account)
	if saveError != nil {
		return nil, saveError
	}

	return &OutputCreateAccountDTO{
		ID: account.ID,
	}, nil
}
