package democrypt

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAESCrypt(t *testing.T) {
	a := assert.New(t)

	_, err := NewAESCrypt(aes128Hash, AesCbcCypher)
	a.NoError(err)
}

func TestAESCryptPadding(t *testing.T) {
	a := assert.New(t)

	aesCrypto, _ := NewAESCrypt(aes128Hash, AesCfbCypher)
	for _, item := range data {
		padded := aesCrypto.pkcs7Padding([]byte(item))
		a.Equal(len(padded)%aes.BlockSize, 0)
		trimmed := aesCrypto.pkcs7Trimming(padded)
		a.Equal(trimmed, []byte(item))
	}
}

func TestAesEncryptDeCypher(t *testing.T) {
	a := assert.New(t)

	for _, Type := range AesChyperModes {
		for _, Hash := range []string{aes128Hash, aes192Hash, aes256Hash} {
			// unavailable cyphers
			if Type == AesCcmCypher || Type == AesCfb8Cypher || Type == AesGcmCypher || Type == AesOfbCypher || Type == AesPcbcCypher {
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
