package democrypt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// RsaCrypt -
type RsaCrypt struct {
	PubKey  *rsa.PublicKey
	PrivKey *rsa.PrivateKey
}

// NewRSACrypt -
func NewRSACrypt(PubKeyPath string, PrivKeyPath string) (*RsaCrypt, error) {
	encTool := &RsaCrypt{}

	if err := encTool.ReadPublicKey(PubKeyPath); err != nil {
		return nil, err
	}
	if err := encTool.ReadPrivateKey(PrivKeyPath); err != nil {
		return nil, err
	}

	return encTool, nil
}

// ReadPublicKey -
func (enc *RsaCrypt) ReadPublicKey(path string) error {
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

// Encrypt will encrypt a string password using an RSA certificate,
// returning a Base64 for of the encrypt result
func (enc *RsaCrypt) Encrypt(password string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, enc.PubKey, []byte(password))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt will decrypt a string password using an RSA certificate,
// expecting a Base64 form of the encrypted password
func (enc *RsaCrypt) Decrypt(password string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.DecryptPKCS1v15(rand.Reader, enc.PrivKey, bytes)
	if err != nil {
		return "", err
	}

	return string(ciphertext), nil
}
