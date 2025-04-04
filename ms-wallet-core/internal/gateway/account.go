package gateway

import "github.com/silverioTenor/fc-wallet-system/internal/entity"

type IAccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}