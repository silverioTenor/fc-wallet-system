package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account, err := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateNewAccountWithInvalidClient(t *testing.T) {
	account, err := NewAccount(nil)

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.Error(t, err, "client is required")
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account, _ := NewAccount(client)

	err := account.Credit(1000)

	assert.Nil(t, err)
	assert.Equal(t, account.Balance, float64(1000))
}

func TestCreditAccountWithInvalidAmount(t *testing.T) {
	client, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account, _ := NewAccount(client)

	err := account.Credit(-100)

	assert.Error(t, err, "amount must be a positive value")
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account, _ := NewAccount(client)

	account.Credit(1000)
	err := account.Debit(200)

	assert.Nil(t, err)
	assert.Equal(t, account.Balance, float64(800))
}

func TestDebitAccountWithInvalidAmount(t *testing.T) {
	client, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account, _ := NewAccount(client)

	err := account.Debit(-100)

	assert.Error(t, err, "amount must be a positive value")
}
