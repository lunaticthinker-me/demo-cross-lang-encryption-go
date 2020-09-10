package democrypt

// Here is a list of demos with more than CBC and CFB modes...
// https://golang.org/src/crypto/cipher/example_test.go

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	// AesTypeCfb - use for AES CFB mode
	AesTypeCfb = iota
	// AesTypeCbc - use for AES CBC mode
	AesTypeCbc
)

// AESCrypt -
type AESCrypt struct {
	Type int
	Key  []byte
	IV   []byte
}

// NewAESCrypt -
func NewAESCrypt(Hash string, Type int) (*AESCrypt, error) {
	enc := AESCrypt{}

	if len(Hash) != 16 && len(Hash) != 24 && len(Hash) != 32 {
		return nil, fmt.Errorf("invalid hash length. must be aes.BlockSize, 24 or 32")
	}

	enc.Type = Type
	enc.Key = []byte(Hash)
	enc.IV = make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, enc.IV); err != nil {
		return nil, fmt.Errorf("unable to generate IV: %s", err.Error())
	}

	return &enc, nil
}

// Encrypt will encrypt a string password using AES algorithm, returning a Base64 for of the encrypt result
func (enc AESCrypt) Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return "", err
	}

	if enc.Type == AesTypeCbc && len(data)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	encdata := make([]byte, len([]byte(data)))

	if enc.Type == AesTypeCbc {
		stream := cipher.NewCBCEncrypter(block, enc.IV)
		stream.CryptBlocks(encdata, []byte(data))
	} else {
		stream := cipher.NewCFBEncrypter(block, enc.IV)
		stream.XORKeyStream(encdata, []byte(data))
	}

	return base64.StdEncoding.EncodeToString(append(enc.IV, encdata...)), nil
}

// Decrypt will decrypt a string password using AES algorithm, expecting a Base64 form of the encrypted password
func (enc AESCrypt) Decrypt(sipher string) (string, error) {

	encsipher, err := base64.StdEncoding.DecodeString(sipher)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return "", err
	}

	enc.IV = encsipher[0:aes.BlockSize]

	encsipher = encsipher[aes.BlockSize:]

	decsipher := make([]byte, len(encsipher))

	if enc.Type == AesTypeCbc {
		stream := cipher.NewCBCDecrypter(block, enc.IV)
		stream.CryptBlocks(decsipher, encsipher)
	} else {
		stream := cipher.NewCFBDecrypter(block, enc.IV)
		stream.XORKeyStream(decsipher, encsipher)
	}

	return string(decsipher), nil
}
