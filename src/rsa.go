package democrypt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

const (
	// PaddingOaep encryption
	PaddingOaep = iota
	// PaddingPkcs1V15 encryption
	PaddingPkcs1V15
)

var RsaPaddings = map[string]int{
	"OAEP":     PaddingOaep,
	"Pkcs1V15": PaddingPkcs1V15,
}

// RsaCrypt -
type RsaCrypt struct {
	Type    int
	PubKey  *rsa.PublicKey
	PrivKey *rsa.PrivateKey
}

// NewRSACrypt -
func NewRSACrypt(pubKeyPath string, privKeyPath string, _type int) (*RsaCrypt, error) {
	encTool := &RsaCrypt{}

	if err := encTool.ReadPublicKey(pubKeyPath); err != nil {
		return nil, err
	}
	if err := encTool.ReadPrivateKey(privKeyPath); err != nil {
		return nil, err
	}

	encTool.Type = _type

	return encTool, nil
}

// ReadPublicKey -
func (enc *RsaCrypt) ReadPublicKey(path string) error {
	// fmt.Println("rsa pub key")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read public key file: %s", err.Error())
	}
	block, _ := pem.Decode(data)
	if block.Type != "PUBLIC KEY" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}
	cert, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	switch cert := cert.(type) {
	case *rsa.PublicKey:
		enc.PubKey = cert
		return nil
	default:
		return fmt.Errorf("found unknown private key type in PKCS#8 wrapping")
	}
}

// ReadPrivateKey -
func (enc *RsaCrypt) ReadPrivateKey(path string) error {
	// fmt.Println("rsa priv key")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read private key file: %s", err.Error())
	}
	block, _ := pem.Decode(data)
	if block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		enc.PrivKey = key
		return nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey:
			enc.PrivKey = key
			return nil
		case *ecdsa.PrivateKey:
			return fmt.Errorf("found ecdsa private key type in PKCS#8 wrapping; aiming for rsa")
		default:
			return fmt.Errorf("found unknown private key type in PKCS#8 wrapping")
		}
	}
	if _, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		// enc.PrivKey = key
		// return nil
		return fmt.Errorf("found ecdsa private key type in PKCS#8 wrapping; aiming for rsa")
	}
	return fmt.Errorf("failed to parse private key")
}

// Decrypt will decrypt a string, using an RSA certificate,
// and expecting a Base64 form of the encrypted string
func (enc *RsaCrypt) Decrypt(ciphertext string) (string, error) {
	plaintext, err := enc.DecryptBytes([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// DecryptBytes will decrypt a set of bytes, using an RSA certificate,
// and expecting a Base64 form of the encrypted set of bytes
func (enc *RsaCrypt) DecryptBytes(ciphertext []byte) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}

	var plaintext []byte
	if enc.Type == PaddingPkcs1V15 {
		plaintext, err = rsa.DecryptPKCS1v15(rand.Reader, enc.PrivKey, bytes)
	} else {
		plaintext, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, enc.PrivKey, bytes, []byte("orders"))
	}
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Encrypt will encrypt a string using an RSA certificate,
// returning a Base64 form of the encrypt result
func (enc *RsaCrypt) Encrypt(plaintext string) (string, error) {
	ciphertext, err := enc.EncryptBytes([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

// EncryptBytes will encrypt a set of bytes using an RSA certificate,
// returning a Base64 form of the encrypt result
var ciphertext []byte

func (enc *RsaCrypt) EncryptBytes(plaintext []byte) ([]byte, error) {
	var err error
	if enc.Type == PaddingPkcs1V15 {
		ciphertext, err = rsa.EncryptPKCS1v15(rand.Reader, enc.PubKey, plaintext)
	} else {
		ciphertext, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, enc.PubKey, plaintext, []byte("orders"))
	}
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(ciphertext)), nil
}
