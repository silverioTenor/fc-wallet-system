package account

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")

	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Willy Wonka", "willy@wonka.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)

	account, _ := entity.NewAccount(s.client)
	s.accountDB.Save(account)

	foundAccount, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, foundAccount.ID)
	s.Equal(account.Balance, foundAccount.Balance)
	s.Equal(account.Client.ID, foundAccount.Client.ID)
	s.Equal(account.Client.Name, foundAccount.Client.Name)
	s.Equal(account.Client.Email, foundAccount.Client.Email)
}
