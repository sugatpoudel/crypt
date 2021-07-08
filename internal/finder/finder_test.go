package finder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tagus/crypt/internal/creds"
)

var (
	ebay = &creds.Credential{
		Id:          "acct-1",
		Service:     "eBay",
		Description: "electronic auction bay",
		Username:    "beanie_babies123",
		Password:    "mars321",
	}
	amazon = &creds.Credential{
		Id:          "acct-2",
		Service:     "Amazon Web Services",
		Description: "america's ali baba",
		Email:       "jeff.bezos@amazon.com",
		Password:    "123jupiter",
	}
	google = &creds.Credential{
		Id:          "acct-3",
		Service:     "Google",
		Description: "big brother",
		Email:       "obrien@google.com",
		Password:    "doublethink",
	}

	crypt = &creds.Crypt{
		Credentials: map[string]*creds.Credential{
			"acct-2": amazon,
			"acct-1": ebay,
			"acct-3": google,
		},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
)

func TestFinder_QueryForNonexistentService(t *testing.T) {
	finder := New(crypt)
	svcs := finder.Filter("non-existent")
	assert.Empty(t, svcs)
}

func TestFinder_QueryByServiceName(t *testing.T) {
	finder := New(crypt)
	svcs := finder.Filter("google")
	assert.Len(t, svcs, 1)
	assert.Equal(t, "acct-3", svcs[0].Id)
	assert.Equal(t, "Google", svcs[0].Service)
}