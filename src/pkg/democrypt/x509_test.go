package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewX509Crypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
	_, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
	a.NoError(err)
}

// func Testx509EncryptDecrypt(t *testing.T) {
// 	a := assert.New(t)

// 	cwd, _ := os.Getwd()
// 	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
// 	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
// 	a.NoError(err)

// 	encrypted, err := X509Crypt.Encrypt("test")
// 	a.NoError(err)

// 	decrypted, err := X509Crypt.Decrypt(encrypted)
// 	a.NoError(err)
// 	a.Equal(decrypted, "test")
// }

// func Testx509Decrypt_FromCSharp(t *testing.T) {
// 	t.Skip("no data available")
// 	a := assert.New(t)

// 	cwd, _ := os.Getwd()
// 	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
// 	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
// 	a.NoError(err)

// 	decrypted, err := X509Crypt.Decrypt("")
// 	a.NoError(err)
// 	a.Equal(decrypted, data)
// }

// func Testx509Decrypt_FromJs(t *testing.T) {
// 	a := assert.New(t)

// 	cwd, _ := os.Getwd()
// 	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
// 	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
// 	a.NoError(err)

// 	decrypted, err := X509Crypt.Decrypt("")
// 	a.NoError(err)
// 	a.Equal(decrypted, data)
// }

// func Testx509Decrypt_FromPy(t *testing.T) {
// 	t.Skip("no data available")
// 	a := assert.New(t)

// 	cwd, _ := os.Getwd()
// 	x509Path := filepath.Join(cwd, "..", "..", "..", "cert", "x509")
// 	X509Crypt, err := NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
// 	a.NoError(err)

// 	decrypted, err := X509Crypt.Decrypt("Hsd/6z3pVxLjPD0wlklknb+EJWC2guY83jo44ItD3IVxrpkgu9Sj1Rdrpq/QHpk3Add3cYaV1sA1eUFQ6jz9c0nTyFNihAG1tp86q9j9+bSk9gVgS8hXmtHTEyjsLo96aNmiyVelB7zR6f9oOGw5gmsxSh+O+EoBAnoPr2+dFiseMKsvRDICqacN+x58DooM4H3Kx7kZRA8nUlnHCTRb8/Odge9J2akBX9MNVhsHoZbDzWEq0CopVSBgOUXpTXwqtxW5bJBgr+hc9w1k8wUZug61J5qI6ASw6iigFdPxoXUzrxGOZa8V3IEbaLGyUli1dWayNgdLToWEtBdRYGysSw==")
// 	a.NoError(err)
// 	a.Equal(decrypted, data)
// }
