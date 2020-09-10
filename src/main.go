package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lunaticthinker-me/demo-cross-lang-encryption-go/src/pkg/democrypt"
)

func main() {
	data := "th1s1smyp@ssw0rd"
	aes128Hash := "1234567890123456"
	aes192Hash := "123456789012345612345678"
	aes256Hash := "12345678901234561234567890123456"

	fmt.Println("AES Encrypted Values:")
	aes, _ := democrypt.NewAESCrypt(aes128Hash, democrypt.AesTypeCfb)
	encrypted, _ := aes.Encrypt(data)
	fmt.Printf("CFB 128 => %s\n", encrypted)
	aes, _ = democrypt.NewAESCrypt(aes192Hash, democrypt.AesTypeCfb)
	encrypted, _ = aes.Encrypt(data)
	fmt.Printf("CFB 192 => %s\n", encrypted)
	aes, _ = democrypt.NewAESCrypt(aes256Hash, democrypt.AesTypeCfb)
	encrypted, _ = aes.Encrypt(data)
	fmt.Printf("CFB 256 => %s\n", encrypted)

	aes, _ = democrypt.NewAESCrypt(aes128Hash, democrypt.AesTypeCbc)
	encrypted, _ = aes.Encrypt(data)
	fmt.Printf("CBC 128 => %s\n", encrypted)
	aes, _ = democrypt.NewAESCrypt(aes192Hash, democrypt.AesTypeCbc)
	encrypted, _ = aes.Encrypt(data)
	fmt.Printf("CBC 192 => %s\n", encrypted)
	aes, _ = democrypt.NewAESCrypt(aes256Hash, democrypt.AesTypeCbc)
	encrypted, _ = aes.Encrypt(data)
	fmt.Printf("CBC 256 => %s\n", encrypted)

	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "cert", "rsa")
	x509Path := filepath.Join(cwd, "cert", "x509")

	fmt.Println("RSA Encrypted Values:")
	rsa, _ := democrypt.NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"))
	encrypted, _ = rsa.Encrypt(data)
	fmt.Println(encrypted)

	fmt.Println("X509 Encrypted Values:")
	x509, _ := democrypt.NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"))
	encrypted, _ = x509.Encrypt(data)
	fmt.Println(encrypted)
}
