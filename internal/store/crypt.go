package store

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/tagus/crypt/internal/creds"
	"github.com/tagus/crypt/internal/secure"
	"golang.org/x/xerrors"
)

const (
	modePerm = 0600
)

// CryptStore represents a crypt instance stored as a file
type CryptStore struct {
	*creds.Crypt
	path   string
	crypto secure.Crypto
}

// createNewStore creates an empty crypt store in the given path
func createNewStore(path string, crypto secure.Crypto, crypt *creds.Crypt) (*CryptStore, error) {
	_, err := os.Stat(path)
	if err == nil {
		return nil, xerrors.New("cryptfile already exists 😬")
	}

	enc, err := crypto.Encrypt(crypt)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(path, enc, modePerm)
	if err != nil {
		return nil, err
	}

	return &CryptStore{
		path:   path,
		crypto: crypto,
		Crypt:  crypt,
	}, nil
}

// InitDefaultStore initializes a defualt crypt store with the given crypt struct
// using the AES crypto implementation. If the crypt file does not exist, one will
// be created in the provided path.
func InitDefaultStore(path, pwd string, crypt *creds.Crypt) (*CryptStore, error) {
	crypto, err := secure.InitAesCrypto(pwd)
	if err != nil {
		return nil, err
	}
	return createNewStore(path, crypto, crypt)
}

// Decrypt attempts to decrypt a crypt store at the given path
func Decrypt(path, pwd string) (*CryptStore, error) {
	crypto, err := secure.InitAesCrypto(pwd)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	crypt, err := crypto.Decrypt(data)
	if err != nil {
		return nil, xerrors.New("password was invalid, decryption failed")
	}

	store := &CryptStore{
		path:   path,
		crypto: crypto,
		Crypt:  crypt,
	}
	return store, nil
}

// Save encrypts the current Crypt and saves it to the path field.
func (s *CryptStore) Save() error {
	s.UpdatedAt = time.Now().Unix()
	data, err := s.crypto.Encrypt(s.Crypt)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.path, data, modePerm)
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
