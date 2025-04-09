package gateway

import "github.com/silverioTenor/fc-wallet-system/internal/entity"

type ITransactionGateway interface {
	Create(transaction *entity.Transaction) error
}