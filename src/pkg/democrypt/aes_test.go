package democrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = "th1s1smyp@ssw0rd"
var aes128Hash = "1234567890123456"
var aes192Hash = "123456789012345612345678"
var aes256Hash = "12345678901234561234567890123456"

func TestNewAESCrypt(t *testing.T) {
	a := assert.New(t)

	_, err := NewAESCrypt(aes128Hash, AesTypeCfb)
	a.NoError(err)
}

// testing CFB

func TestAESEncryptDecrypt_128_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes128Hash, AesTypeCfb)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESEncryptDecrypt_192_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes192Hash, AesTypeCfb)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESEncryptDecrypt_256_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes256Hash, AesTypeCfb)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

// testing CBC

func TestAESEncryptDecrypt_128_Cbc(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes128Hash, AesTypeCbc)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESEncryptDecrypt_192_Cbc(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes192Hash, AesTypeCbc)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESEncryptDecrypt_256_Cbc(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes256Hash, AesTypeCbc)

	encrypted, err := aes.Encrypt(data)
	a.NoError(err)

	decrypted, err := aes.Decrypt(encrypted)
	a.NoError(err)

	a.Equal(decrypted, data)
}

// testing JS CFB

func TestAESDecrypt_FromJS_128_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes128Hash, AesTypeCfb)

	decrypted, err := aes.Decrypt("GigmhLJurG5BhbKZ/4Rbh52d+uv8HoBBFl55d0QzKNQ=")
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESDecrypt_FromJS_192_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes192Hash, AesTypeCfb)

	decrypted, err := aes.Decrypt("wTbvR63MrpuZmfMS0//0nbdQseZnSC2vm61/rUPVbnc=")
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESDecrypt_FromJS_256_Cfb(t *testing.T) {
	a := assert.New(t)

	aes, err := NewAESCrypt(aes256Hash, AesTypeCfb)

	decrypted, err := aes.Decrypt("iDXB861O/V3Cfn9Mexrn1a4hOYxmXrsxcWExy+utt+o=")
	a.NoError(err)

	a.Equal(decrypted, data)
}

// testing JS CBC

func TestAESDecrypt_FromJS_128_Cbc(t *testing.T) {
	t.Skip("algorithm missmatching")
	a := assert.New(t)

	aes, err := NewAESCrypt(aes128Hash, AesTypeCbc)

	decrypted, err := aes.Decrypt("hsub5jmTAFsBBy/1GyNZi/tI35lJX6u4E1BuhoD5t/NH8lvUJarhVE8lCIkiF4g2")
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESDecrypt_FromJS_192_Cbc(t *testing.T) {
	t.Skip("algorithm missmatching")
	a := assert.New(t)

	aes, err := NewAESCrypt(aes192Hash, AesTypeCbc)

	decrypted, err := aes.Decrypt("wLzKcST0ltRik0/lB4+wTz0BOQLamK9ZXtevtP89mO5IIXKg3kxJvPvyroFw4F05")
	a.NoError(err)

	a.Equal(decrypted, data)
}

func TestAESDecrypt_FromJS_256_Cbc(t *testing.T) {
	t.Skip("algorithm missmatching")
	a := assert.New(t)

	aes, err := NewAESCrypt(aes256Hash, AesTypeCbc)

	decrypted, err := aes.Decrypt("KO2RwoX8cd3ydOe5E6/kIAB2o40ILhBLjkToAy8pLRXhNj25cFx0bSO5ThXHkTpN")
	a.NoError(err)

	a.Equal(decrypted, data)
}
