package client

import (
	"time"

	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/silverioTenor/fc-wallet-system/internal/gateway"
)

type InputCreateClientDTO struct {
	Name  string
	Email string
}

type OutputCreateClientDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.IClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.IClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (usecase *CreateClientUseCase) Execute(
	input InputCreateClientDTO,
) (*OutputCreateClientDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = usecase.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &OutputCreateClientDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
