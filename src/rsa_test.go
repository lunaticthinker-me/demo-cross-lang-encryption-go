package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRsa(Type int) (*RsaCrypt, error) {
	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "cert", "rsa")
	return NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"), Type)
}

func TestNewRSACrypt(t *testing.T) {
	a := assert.New(t)

	_, err := getRsa(PaddingOaep)
	a.NoError(err)
}

func TestRSAEncryptDecrypt_Oaep(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(PaddingOaep)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := rsacrypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := rsacrypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestRSAEncryptDecrypt_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(PaddingPkcs1V15)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := rsacrypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := rsacrypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}
