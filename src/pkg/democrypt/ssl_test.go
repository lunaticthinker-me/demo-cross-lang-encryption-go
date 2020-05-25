package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSSLCrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	sslPath := filepath.Join(cwd, "..", "..", "..", "cert", "ssl")
	_, err := NewSSLCrypt(filepath.Join(sslPath, "cert.pem"), filepath.Join(sslPath, "key.pem"))
	a.NoError(err)
}

func TestSSLEncrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	sslPath := filepath.Join(cwd, "..", "..", "..", "cert", "ssl")
	sslcrypt, err := NewSSLCrypt(filepath.Join(sslPath, "cert.pem"), filepath.Join(sslPath, "key.pem"))
	a.NoError(err)

	_, err = sslcrypt.Encrypt("test")
	a.NoError(err)
}

func TestSSLDecrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	sslPath := filepath.Join(cwd, "..", "..", "..", "cert", "ssl")
	sslcrypt, err := NewSSLCrypt(filepath.Join(sslPath, "cert.pem"), filepath.Join(sslPath, "key.pem"))
	a.NoError(err)

	encrypted, err := sslcrypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := sslcrypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}
