package democrypt

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// X509Crypt -
type X509Crypt struct {
	RsaCrypt
}

// NewX509Crypt -
func NewX509Crypt(PubKeyPath string, PrivKeyPath string, Type int) (*X509Crypt, error) {
	encTool := &X509Crypt{}

	if err := encTool.ReadPublicKey(PubKeyPath); err != nil {
		return nil, err
	}
	if err := encTool.ReadPrivateKey(PrivKeyPath); err != nil {
		return nil, err
	}

	encTool.Type = Type

	return encTool, nil
}

// ReadPublicKey -
func (enc *X509Crypt) ReadPublicKey(path string) error {
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
func (enc *X509Crypt) ReadPrivateKey(path string) error {
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
