package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Willy Wonka", "willy@wonka.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Willy Wonka", client.Name)
	assert.Equal(t, "willy@wonka.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}