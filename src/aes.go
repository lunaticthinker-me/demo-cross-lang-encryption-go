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

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

const (
	AesCbcCypher = iota
	AesCcmCypher
	AesCfbCypher
	AesCfb8Cypher
	AesCtrCypher
	AesEcbCypher
	AesGcmCypher
	AesOfbCypher
	AesPcbcCypher
)

var AesChyperModes = map[string]int{
	"CBC":  AesCbcCypher,  // ✓
	"CCM":  AesCcmCypher,  // not implemented https://github.com/pschlump/AesCCM/blob/master/ccm_test.go
	"CFB":  AesCfbCypher,  // ✓
	"CFB8": AesCfb8Cypher, // unavailable
	"CTR":  AesCtrCypher,  // ✓
	"ECB":  AesEcbCypher,  // ✓
	"GCM":  AesGcmCypher,  // see AesCcmCypher
	"OFB":  AesOfbCypher,  // unavailable
	"PCBC": AesPcbcCypher, // unavailable
}

// AESCrypt -
type AESCrypt struct {
	Type int
	Key  []byte
	IV   []byte
}

func newDecryptEntity(Type int) {}

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
	switch enc.Type {
	case AesCbcCypher:
		paddedtext := enc.pkcs7Padding(plaintext)
		ciphertext = make([]byte, len(paddedtext))

		mode := cipher.NewCBCEncrypter(block, enc.IV)
		mode.CryptBlocks(ciphertext, paddedtext)
	case AesCfbCypher:
		ciphertext = make([]byte, len(plaintext))

		mode := cipher.NewCFBEncrypter(block, enc.IV)
		mode.XORKeyStream(ciphertext, plaintext)
	case AesCtrCypher:
		ciphertext = make([]byte, len(plaintext))

		mode := cipher.NewCTR(block, enc.IV)
		mode.XORKeyStream(ciphertext, plaintext)
	case AesEcbCypher:
		mode := ecb.NewECBEncrypter(block)
		plaintext, err = padding.NewPkcs7Padding(mode.BlockSize()).Pad(plaintext)
		if err != nil {
			return nil, err
		}

		ciphertext = make([]byte, len(plaintext))
		mode.CryptBlocks(ciphertext, plaintext)
	default:
		for cypherLabel, cypherType := range AesChyperModes {
			if cypherType == enc.Type {
				return nil, fmt.Errorf("invalid cipher type: %s", cypherLabel)
			}
		}
		return nil, fmt.Errorf("invalid cipher type int value")
	}

	cipher := base64.StdEncoding.EncodeToString(append(enc.IV, ciphertext...))
	// fmt.Printf("\n%v => IV:%d CIPHER:%d\n", cipher, len(enc.IV), len(ciphertext))
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

	// fmt.Printf("\n%v => %d (%d)\n", encrypted, len(encrypted), len(decoded))
	decrypted := make([]byte, len(encrypted))

	switch enc.Type {
	case AesCbcCypher:
		stream := cipher.NewCBCDecrypter(block, enc.IV)
		stream.CryptBlocks(decrypted, encrypted)
		decrypted = enc.pkcs7Trimming(decrypted)
	case AesCfbCypher:
		stream := cipher.NewCFBDecrypter(block, enc.IV)
		stream.XORKeyStream(decrypted, encrypted)
	case AesCtrCypher:
		stream := cipher.NewCTR(block, enc.IV)
		stream.XORKeyStream(decrypted, encrypted)
	case AesEcbCypher:
		mode := ecb.NewECBDecrypter(block)
		mode.CryptBlocks(decrypted, encrypted)

		decrypted, err = padding.NewPkcs7Padding(mode.BlockSize()).Unpad(decrypted)
		if err != nil {
			return nil, err
		}
	default:
		for cypherLabel, cypherType := range AesChyperModes {
			if cypherType == enc.Type {
				return nil, fmt.Errorf("invalid cipher type: %s", cypherLabel)
			}
		}
		return nil, fmt.Errorf("invalid cipher type int value")
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
