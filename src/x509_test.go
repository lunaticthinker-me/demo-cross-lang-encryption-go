package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getX509(Type int) (*X509Crypt, error) {
	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "cert", "x509")
	return NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"), Type)
}

func TestNewX509Crypt(t *testing.T) {
	a := assert.New(t)

	_, err := getX509(PaddingOaep)
	a.NoError(err)
}

func TestX509EncryptDecrypt_Oaep(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getX509(PaddingOaep)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := x509crypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := x509crypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestX509EncryptDecrypt_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getX509(PaddingPkcs1V15)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := x509crypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := x509crypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}
