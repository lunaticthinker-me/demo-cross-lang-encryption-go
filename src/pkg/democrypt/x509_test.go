package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewX509Crypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
	_, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
	a.NoError(err)
}

func Testx509Encrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
	a.NoError(err)

	_, err = X509Crypt.Encrypt("test")
	a.NoError(err)
}

func Testx509Decrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
	a.NoError(err)

	encrypted, err := X509Crypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := X509Crypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}
