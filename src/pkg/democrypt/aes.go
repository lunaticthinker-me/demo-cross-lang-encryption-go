package democrypt

// Here is a list of demos with more than CBC and CFB modes...
// https://golang.org/src/crypto/cipher/example_test.go

import (
	"bytes"
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

// Encrypt will encrypt a string using AES algorithm, returning a Base64 form of the result
func (enc AESCrypt) Encrypt(plaintext string) (string, error) {
	cipherbytes, err := enc.EncryptBytes([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return string(cipherbytes), nil
}

// EncryptBytes will encrypt a set of bytes using AES algorithm, returning a Base64 form of the result
func (enc AESCrypt) EncryptBytes(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return nil, err
	}

	var ciphertext []byte
	if enc.Type == AesTypeCbc {
		paddedtext := enc.pkcs7Padding(plaintext)
		ciphertext = make([]byte, len(paddedtext))

		aesCbc := cipher.NewCBCEncrypter(block, enc.IV)
		aesCbc.CryptBlocks(ciphertext, paddedtext)
	} else {
		ciphertext = make([]byte, len(plaintext))

		aesCfb := cipher.NewCFBEncrypter(block, enc.IV)
		aesCfb.XORKeyStream(ciphertext, plaintext)
	}

	cipher := base64.StdEncoding.EncodeToString(append(enc.IV, ciphertext...))
	return []byte(cipher), nil
}

// Decrypt will decrypt a Base64 encoded string using AES algorithm returning a string
func (enc AESCrypt) Decrypt(ciphertext string) (string, error) {
	decrypted, err := enc.DecryptBytes([]byte(ciphertext))
	return string(decrypted), err
}

// DecryptBytes will decrypt a set of Base64 encoded bytes using AES algorithm returning a byte array
func (enc AESCrypt) DecryptBytes(cipherbytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return nil, err
	}

	decodedStr, err := base64.StdEncoding.DecodeString(string(cipherbytes))
	if err != nil {
		return nil, err
	}
	decoded := []byte(decodedStr)

	enc.IV = decoded[0:aes.BlockSize]
	encrypted := decoded[aes.BlockSize:]

	fmt.Printf("\n%v => %d\n", encrypted, len(encrypted))
	decrypted := make([]byte, len(encrypted))

	if enc.Type == AesTypeCbc {
		stream := cipher.NewCBCDecrypter(block, enc.IV)
		stream.CryptBlocks(decrypted, encrypted)
		decrypted = enc.pkcs7Trimming(decrypted)
	} else {
		stream := cipher.NewCFBDecrypter(block, enc.IV)
		stream.XORKeyStream(decrypted, encrypted)
	}

	return decrypted, nil
}

// https://github.com/mervick/aes-everywhere
func (enc AESCrypt) pkcs7Padding(ciphertext []byte) []byte {
	bs := aes.BlockSize
	padding := bs - len(ciphertext)%bs
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (enc AESCrypt) pkcs7Trimming(plaintext []byte) []byte {
	padding := plaintext[len(plaintext)-1]
	trimmed := plaintext[:len(plaintext)-int(padding)]
	return trimmed
}
