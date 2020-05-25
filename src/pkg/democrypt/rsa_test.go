package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRSACrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	_, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)
}

func TestRSAEncrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	_, err = rsacrypt.Encrypt("test")
	a.NoError(err)
}

func TestRSADecrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	encrypted, err := rsacrypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}
