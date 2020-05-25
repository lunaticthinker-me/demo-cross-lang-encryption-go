package democrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AESCrypt -
type AESCrypt struct {
	Key []byte
	IV  []byte
}

// NewAESCrypt -
func NewAESCrypt(Hash string) (*AESCrypt, error) {
	enc := AESCrypt{}

	if len(Hash) != 16 && len(Hash) != 24 && len(Hash) != 32 {
		return nil, fmt.Errorf("invalid hash length. must be 16, 24 or 32")
	}

	enc.Key = []byte(Hash)
	enc.IV = make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, enc.IV); err != nil {
		return nil, fmt.Errorf("unable to generate IV: %s", err.Error())
	}

	return &enc, nil
}

// Encrypt will encrypt a string password using AES algorithm, returning a Base64 for of the encrypt result
func (enc AESCrypt) Encrypt(password string) (string, error) {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, enc.IV)

	encPassword := make([]byte, len([]byte(password)))

	stream.XORKeyStream(encPassword, []byte(password))

	return base64.StdEncoding.EncodeToString(append(enc.IV, encPassword...)), nil
}

// Decrypt will decrypt a string password using AES algorithm, expecting a Base64 form of the encrypted password
func (enc AESCrypt) Decrypt(password string) (string, error) {

	encPassword, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return "", err
	}

	enc.IV = encPassword[0:16]

	stream := cipher.NewCFBDecrypter(block, enc.IV)

	encPassword = encPassword[16:]

	decPassword := make([]byte, len(encPassword))

	stream.XORKeyStream(decPassword, encPassword)

	return string(decPassword), nil
}
