package transaction

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id, account_from_id varchar(255), account_to_id varchar(255), amount float, created_at date)")

	s.clientFrom, _ = entity.NewClient("Willy Wonka", "willy@wonka.com")
	s.accountFrom, _ = entity.NewAccount(s.clientFrom)
	s.accountFrom.Balance = 1000

	s.clientTo, _ = entity.NewClient("Jane Wonka", "jane@wonka.com")
	s.accountTo, _ = entity.NewAccount(s.clientTo)
	s.accountTo.Balance = 1000

	s.transactionDB = NewTransactionDB(db)
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}
func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 200)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
