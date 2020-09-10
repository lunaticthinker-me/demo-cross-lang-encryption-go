package democrypt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRSACrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	_, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)
}

func TestRSAEncryptDecrypt(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	encrypted, err := rsacrypt.Encrypt("test")
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt(encrypted)
	a.NoError(err)
	a.Equal(decrypted, "test")
}

func TestRSADecrypt_FromCSharp(t *testing.T) {
	t.Skip("no data available")
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt("")
	a.NoError(err)
	a.Equal(decrypted, data)
}

func TestRSADecrypt_FromJs(t *testing.T) {
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt("HDIXtNnbJjaxfCKFEmecq2LL32nDmzn+p5FsZZZYtDUKrf6VPAFI2iNxXVlYcJjoJ41z70zRQg1+Y28bfAvlplqCRV1REXDsVMWWTq02g/oJTd2bwB/0ly3lb6tEFbCK6JQ9WBLU4LcX6qgZsJZWJIWhBimNFA+wEIwFxZcR5csqg3/YDdnVafD8ydywsQPp7pQ3FbFseHt15nNJ9Hg+jbbjb3n9TAeJ0J8BJ0q7ocd9tbC46VaviaAI2tyF9sarMdR2AdywwrLzCaJk1a1UibCCvVH58H+08utb/RjvWj/VU1UfDdvmeoinxQKLVbzZDfkGmctDE1AcOW4CJ0+Gzw==")
	a.NoError(err)
	a.Equal(decrypted, data)
}

func TestRSADecrypt_FromPy(t *testing.T) {
	t.Skip("no data available")
	a := assert.New(t)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "..", "..", "..", "cert", "rsa")
	rsacrypt, err := NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	a.NoError(err)

	decrypted, err := rsacrypt.Decrypt("")
	a.NoError(err)
	a.Equal(decrypted, data)
}
