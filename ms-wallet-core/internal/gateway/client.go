package gateway

import "github.com/silverioTenor/fc-wallet-system/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}