package creds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ebay = Credential{
		Service:     "eBay",
		Description: "electronic auction bay",
		Username:    "beanie_babies123",
		Password:    "mars321",
	}
	amazon = Credential{
		Service:     "Amazon Web Services",
		Description: "america's ali baba",
		Email:       "jeff.bezos@amazon.com",
		Password:    "123jupiter",
	}
	crypt = &Crypt{
		Credentials: map[string]Credential{
			"ebay": ebay,
		},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
)

func TestSetCredential(t *testing.T) {
	crypt.SetCredential(amazon)
	cred := crypt.FindCredential("Amazon Web    Services")
	assert.NotNil(t, cred, "Did not set credential properly")
}

func TestRemoveCredential(t *testing.T) {
	crypt.RemoveCredential("ebay")
	cred := crypt.FindCredential("ebay")
	assert.Nil(t, cred, "Did not remove credential properly")
}

func TestGetJSON(t *testing.T) {
	_, err := crypt.GetJSON()
	assert.Nil(t, err, "Did not marshall properly")
}

func TestFindCredential(t *testing.T) {
	cred := crypt.FindCredential("amazon web services")
	assert.NotNil(t, cred, "Did not retrieve credential")
	assert.Equal(t, "Amazon Web Services", cred.Service)
}
