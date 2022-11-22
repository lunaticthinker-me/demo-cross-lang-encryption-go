package democrypt

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAESCrypt(t *testing.T) {
	a := assert.New(t)

	_, err := NewAESCrypt(aes128Hash, AesCbcCrypt)
	a.NoError(err)
}

func TestAESCryptPadding(t *testing.T) {
	a := assert.New(t)

	aesCrypto, _ := NewAESCrypt(aes128Hash, AesCfbCrypt)
	for _, item := range data {
		padded := aesCrypto.pkcs7Padding([]byte(item))
		a.Equal(len(padded)%aes.BlockSize, 0)
		trimmed := aesCrypto.pkcs7Trimming(padded)
		a.Equal(trimmed, []byte(item))
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	a := assert.New(t)

	for _, Type := range AesCryptModes {
		for _, Hash := range []string{aes128Hash, aes192Hash, aes256Hash} {
			if Type == AesCcmCrypt || Type == AesCfb8Crypt || Type == AesGcmCrypt
        || Type == AesOfbCrypt || Type == AesPcbcCrypt {
				continue
			}
			aesCrypto, _ := NewAESCrypt(Hash, Type)
			for _, item := range data {
				encrypted, err := aesCrypto.Encrypt(item)
				a.NoError(err)

				decrypted, err := aesCrypto.Decrypt(encrypted)
				a.NoError(err)
				a.Equal(item, decrypted)
			}
		}
	}

}
