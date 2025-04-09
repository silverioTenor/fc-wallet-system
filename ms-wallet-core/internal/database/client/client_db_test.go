package client

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-wallet-system/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	suite.clientDB = NewClientDB(db)
}

func (suite *ClientDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID: "1",
		Name: "Willy Wonka",
		Email: "willy@wonka.com",
	}
	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Willy", "willy@wonka.com")
	suite.clientDB.Save(client)

	foundClient, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, foundClient.ID)
	suite.Equal(client.Name, foundClient.Name)
	suite.Equal(client.Email, foundClient.Email)
}
