package democrypt

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// aes192Cfb, _ := NewAESCrypt(aes192Hash, AesTypeCfb)
// aes256Cfb, _ := NewAESCrypt(aes256Hash, AesTypeCfb)
// aes128Cbc, _ := NewAESCrypt(aes128Hash, AesTypeCbc)
// aes192Cbc, _ := NewAESCrypt(aes192Hash, AesTypeCbc)
// aes256Cbc, _ := NewAESCrypt(aes256Hash, AesTypeCbc)

func TestNewAESCrypt(t *testing.T) {
	a := assert.New(t)

	_, err := NewAESCrypt(aes128Hash, AesTypeCfb)
	a.NoError(err)
}

func TestAESCryptPadding(t *testing.T) {
	a := assert.New(t)

	aes128Cfb, _ := NewAESCrypt(aes128Hash, AesTypeCfb)
	for _, item := range data {
		padded := aes128Cfb.pkcs7Padding([]byte(item))
		a.Equal(len(padded)%aes.BlockSize, 0)
		trimmed := aes128Cfb.pkcs7Trimming(padded)
		a.Equal(trimmed, []byte(item))
	}
}

// testing CFB

func TestAESEncryptDecrypt_128_Cfb(t *testing.T) {
	a := assert.New(t)

	aes128Cfb, _ := NewAESCrypt(aes128Hash, AesTypeCfb)
	for _, item := range data {
		encrypted, err := aes128Cfb.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes128Cfb.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestAESEncryptDecrypt_192_Cfb(t *testing.T) {
	a := assert.New(t)

	aes192Cfb, _ := NewAESCrypt(aes192Hash, AesTypeCfb)
	for _, item := range data {
		encrypted, err := aes192Cfb.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes192Cfb.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestAESEncryptDecrypt_256_Cfb(t *testing.T) {
	a := assert.New(t)

	aes256Cfb, _ := NewAESCrypt(aes256Hash, AesTypeCfb)
	for _, item := range data {
		encrypted, err := aes256Cfb.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes256Cfb.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

// testing CBC

func TestAESEncryptDecrypt_128_Cbc(t *testing.T) {
	a := assert.New(t)

	aes128Cbc, _ := NewAESCrypt(aes128Hash, AesTypeCbc)
	for _, item := range data {
		encrypted, err := aes128Cbc.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes128Cbc.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestAESEncryptDecrypt_192_Cbc(t *testing.T) {
	a := assert.New(t)

	aes192Cbc, _ := NewAESCrypt(aes192Hash, AesTypeCbc)
	for _, item := range data {
		encrypted, err := aes192Cbc.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes192Cbc.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

func TestAESEncryptDecrypt_256_Cbc(t *testing.T) {
	a := assert.New(t)

	aes256Cbc, _ := NewAESCrypt(aes256Hash, AesTypeCbc)
	for _, item := range data {
		encrypted, err := aes256Cbc.Encrypt(item)
		a.NoError(err)

		decrypted, err := aes256Cbc.Decrypt(encrypted)
		a.NoError(err)
		a.Equal(decrypted, item)
	}
}

// testing CS CFB

func TestAESDecrypt_From_CS_128_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCfb)

	decrypted, err := aes128Cbf.Decrypt(CS_AES_CFB8_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_CS_192_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCfb)

	decrypted, err := aes192Cbf.Decrypt(CS_AES_CFB8_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_CS_256_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCfb)

	decrypted, err := aes256Cbf.Decrypt(CS_AES_CFB8_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

// testing CS CFB

func TestAESDecrypt_From_CS_128_Cbc(t *testing.T) {
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCbc)

	decrypted, err := aes128Cbf.Decrypt(CS_AES_CBC_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_CS_192_Cbc(t *testing.T) {
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCbc)

	decrypted, err := aes192Cbf.Decrypt(CS_AES_CBC_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_CS_256_Cbc(t *testing.T) {
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCbc)

	decrypted, err := aes256Cbf.Decrypt(CS_AES_CBC_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

// testing JS CFB

func TestAESDecrypt_From_JS_128_Cfb(t *testing.T) {
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCfb)

	decrypted, err := aes128Cbf.Decrypt(JS_AES_CFB_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_JS_192_Cfb(t *testing.T) {
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCfb)

	decrypted, err := aes192Cbf.Decrypt(JS_AES_CFB_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_JS_256_Cfb(t *testing.T) {
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCfb)

	decrypted, err := aes256Cbf.Decrypt(JS_AES_CFB_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

// testing JS CFB

func TestAESDecrypt_From_JS_128_Cbc(t *testing.T) {
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCbc)

	decrypted, err := aes128Cbf.Decrypt(JS_AES_CBC_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_JS_192_Cbc(t *testing.T) {
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCbc)

	decrypted, err := aes192Cbf.Decrypt(JS_AES_CBC_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_JS_256_Cbc(t *testing.T) {
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCbc)

	decrypted, err := aes256Cbf.Decrypt(JS_AES_CBC_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

// testing PY CFB

func TestAESDecrypt_From_Py_128_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCfb)

	decrypted, err := aes128Cbf.Decrypt(PY_AES_CFB8_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_Py_192_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCfb)

	decrypted, err := aes192Cbf.Decrypt(PY_AES_CFB8_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_Py_256_Cfb(t *testing.T) {
	t.Skip("golang does not support cfb8")
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCfb)

	decrypted, err := aes256Cbf.Decrypt(PY_AES_CFB8_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

// testing PY CFB

func TestAESDecrypt_From_Py_128_Cbc(t *testing.T) {
	a := assert.New(t)

	aes128Cbf, err := NewAESCrypt(aes128Hash, AesTypeCbc)

	decrypted, err := aes128Cbf.Decrypt(PY_AES_CBC_128)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_Py_192_Cbc(t *testing.T) {
	a := assert.New(t)

	aes192Cbf, err := NewAESCrypt(aes192Hash, AesTypeCbc)

	decrypted, err := aes192Cbf.Decrypt(PY_AES_CBC_192)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}

func TestAESDecrypt_From_Py_256_Cbc(t *testing.T) {
	a := assert.New(t)

	aes256Cbf, err := NewAESCrypt(aes256Hash, AesTypeCbc)

	decrypted, err := aes256Cbf.Decrypt(PY_AES_CBC_256)
	a.NoError(err)
	a.Equal(decrypted, data[0])
}
