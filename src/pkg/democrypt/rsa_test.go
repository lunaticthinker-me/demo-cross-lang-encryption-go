package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRsa(Type int) (*RsaCrypt, error) {
	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	return NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"), Type)
}

func TestNewRSACrypt(t *testing.T) {
	a := assert.New(t)

	_, err := getRsa(RsaOaep)
	a.NoError(err)
}

func TestRSAEncryptDecrypt_Oaep(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaOaep)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := rsacrypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := rsacrypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestRSAEncryptDecrypt_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	for _, item := range data {
		encrypted, err := rsacrypt.Encrypt(item)
		a.NoError(err)

		decrypted, err := rsacrypt.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestRSADecrypt_FromCSharp_Oaep(t *testing.T) {
	skipNotCompatible(t)
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaOaep)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(CS_RSA_OAEP)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestRSADecrypt_FromCSharp_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(CS_RSA_PKCS1V1_5)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestRSADecrypt_FromJs_Oaep(t *testing.T) {
	skipNotCompatible(t)
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaOaep)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(JS_RSA_OAEP)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestRSADecrypt_FromJs_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(JS_RSA_PKCS1V1_5)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestRSADecrypt_FromPy_Oaep(t *testing.T) {
	skipNotCompatible(t)
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaOaep)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(PY_RSA_OAEP)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestRSADecrypt_FromPy_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	rsacrypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(PY_RSA_PKCS1V1_5)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}
