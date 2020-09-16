package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getX509(Type int) (*X509Crypt, error) {
	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
	return NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"), Type)
}

func TestNewX509Crypt(t *testing.T) {
	a := assert.New(t)

	_, err := getRsa(RsaOaep)
	a.NoError(err)
}

func Testx509EncryptDecrypt_Oaep(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaOaep)
	a.NoError(err)

	encrypted, err := x509crypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}

func Testx509EncryptDecrypt_Pkcs1v15(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	encrypted, err := x509crypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}

// func Testx509Decrypt_FromCs_Oaep(t *testing.T) {
// 	a := assert.New(t)

// 	x509crypt, err := getRsa(RsaOaep)
// 	a.NoError(err)

// 	decrypted, err := x509crypt.Decrypt(CS_X509_OAEP)
// 	a.NoError(err)
// 	a.Equal(decrypted, data[0])
// }

// func Testx509Decrypt_FromCs_Pkcs1V15(t *testing.T) {
// 	a := assert.New(t)

// 	x509crypt, err := getRsa(RsaPkcs1V15)
// 	a.NoError(err)

// 	decrypted, err := x509crypt.Decrypt(CS_X509_PKCS1V1_5)
// 	a.NoError(err)
// 	a.Equal(decrypted, data[0])
// }

func Testx509Decrypt_FromJs_Oaep(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaOaep)
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(JS_X509_OAEP)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func Testx509Decrypt_FromJs_Pkcs1V15(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(JS_X509_PKCS1V1_5)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func Testx509Decrypt_FromPy_Oaep(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaOaep)
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(PY_X509_OAEP)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func Testx509Decrypt_FromPy_Pkcs1V15(t *testing.T) {
	a := assert.New(t)

	x509crypt, err := getRsa(RsaPkcs1V15)
	a.NoError(err)

	decrypted, err := x509crypt.Decrypt(PY_X509_PKCS1V1_5)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}
