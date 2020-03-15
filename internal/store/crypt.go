package store

import (
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/sugatpoudel/crypt/internal/creds"
	"github.com/sugatpoudel/crypt/internal/secure"
)

const perm = 0600

// CryptStore represents a crypt instance stored as a file
type CryptStore struct {
	path   string
	crypto secure.Crypto
	Crypt  *creds.Crypt
}

// Creates an empty crypt file in the given path.
func createDefaultCryptFile(path string, crypto secure.Crypto) error {
	credMap := make(map[string]creds.Credential)
	now := time.Now().Unix()
	crypt := &creds.Crypt{
		Credentials: credMap,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	enc, err := crypto.Encrypt(crypt)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, enc, perm)
	if err != nil {
		return err
	}

	return nil
}

// InitDefaultStore initializes a default crypt store using the AES crypto implementation.
// If the crypt file does not exist, one will be created in the provided path.
func InitDefaultStore(path, pwd string) (*CryptStore, error) {
	crypto, err := secure.InitAesCrypto(pwd)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := createDefaultCryptFile(path, crypto)
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	crypt, err := crypto.Decrypt(data)
	if err != nil {
		return nil, errors.New("password was invalid, decryption failed")
	}

	store := &CryptStore{path, crypto, crypt}
	return store, nil
}

// Save encrypts the current Crypt and saves it to the path field.
func (s *CryptStore) Save() error {
	s.Crypt.UpdatedAt = time.Now().Unix()
	data, err := s.crypto.Encrypt(s.Crypt)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.path, data, perm)
	if err != nil {
		return err
	}

	return nil
}

// ChangePwd recreates the Crypto instance with the new password.
func (s *CryptStore) ChangePwd(pwd string) error {
	crypto, err := secure.InitAesCrypto(pwd)
	if err != nil {
		return err
	}

	s.crypto = crypto
	return nil
}