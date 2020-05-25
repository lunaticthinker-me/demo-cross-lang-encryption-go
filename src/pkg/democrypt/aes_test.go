package democrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAESCrypt(t *testing.T) {
	a := assert.New(t)

	_, err := NewAESCrypt("1234567890123456")
	a.NoError(err)
}

func TestAESEncrypt(t *testing.T) {
	a := assert.New(t)

	aescryp, err := NewAESCrypt("1234567890123456")

	_, err = aescryp.Encrypt("test")
	a.NoError(err)
}

func TestAESDecrypt(t *testing.T) {
	a := assert.New(t)

	aescryp, err := NewAESCrypt("1234567890123456")

	encrypted, err := aescryp.Encrypt("test")
	a.NoError(err)

	decrypted, err := aescryp.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}
