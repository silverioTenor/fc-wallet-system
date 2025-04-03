package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Madre Tereza de Calcutá", "madre.tereza@calcuta.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 500)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(500), account1.Balance)
	assert.Equal(t, float64(1500), account2.Balance)
}

func TestCreateTransactionWithAmountGreaterThanBalance(t *testing.T) {
	client1, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Madre Tereza de Calcutá", "madre.tereza@calcuta.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 2000)

	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Equal(t, float64(1000), account1.Balance)
	assert.Equal(t, float64(1000), account2.Balance)
}

func TestCreateTransactionWithInvalidAmount(t *testing.T) {
	client1, _ := NewClient("Willy Wonka", "willy@wonka.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Madre Tereza de Calcutá", "madre.tereza@calcuta.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -2000)

	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be grater than zero")
	assert.Equal(t, float64(1000), account1.Balance)
	assert.Equal(t, float64(1000), account2.Balance)
}