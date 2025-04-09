package transaction

import (
	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/silverioTenor/fc-wallet-system/internal/gateway"
)

type InputCreateTransactionDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type OutputCreateTransactionDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.ITransactionGateway
	AccountGateway     gateway.IAccountGateway
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.ITransactionGateway,
	accountGateway gateway.IAccountGateway,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (usecase *CreateTransactionUseCase) Execute(
	input InputCreateTransactionDTO,
) (*OutputCreateTransactionDTO, error) {
	accountFrom, err := usecase.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := usecase.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = usecase.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &OutputCreateTransactionDTO{ID: transaction.ID}, nil
}
