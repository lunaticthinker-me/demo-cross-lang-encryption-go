package democrypt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// SslCrypt -
type SslCrypt struct {
	PubKey  *rsa.PublicKey
	PrivKey *rsa.PrivateKey
}

// NewSSLCrypt -
func NewSSLCrypt(PubKeyPath string, PrivKeyPath string) (*SslCrypt, error) {
	encTool := &SslCrypt{}

	if err := encTool.ReadPublicKey(PubKeyPath); err != nil {
		return nil, err
	}
	if err := encTool.ReadPrivateKey(PrivKeyPath); err != nil {
		return nil, err
	}

	return encTool, nil
}

// ReadPublicKey -
func (enc *SslCrypt) ReadPublicKey(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read public key file: %s", err.Error())
	}

	block, _ := pem.Decode(data)
	if block.Type != "CERTIFICATE" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	enc.PubKey = cert.PublicKey.(*rsa.PublicKey)

	return nil
}

// ReadPrivateKey -
func (enc *SslCrypt) ReadPrivateKey(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read private key file: %s", err.Error())
	}

	block, _ := pem.Decode(data)
	if block.Type != "PRIVATE KEY" {
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

// Encrypt will encrypt a string password using an SSL certificate,
// returning a Base64 for of the encrypt result
func (enc *SslCrypt) Encrypt(password string) (string, error) {
	hash := sha512.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, enc.PubKey, []byte(password), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt will decrypt a string password using an SSL certificate,
// expecting a Base64 form of the encrypted password
func (enc *SslCrypt) Decrypt(password string) (string, error) {
	hash := sha512.New()

	bytes, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.DecryptOAEP(hash, rand.Reader, enc.PrivKey, bytes, nil)
	if err != nil {
		return "", err
	}

	return string(ciphertext), nil
}
