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

	var terms [][]string
	terms = append(terms, []string{"C#", "public static string GO", "\""})
	terms = append(terms, []string{"Js", "export const GO", "'"})
	terms = append(terms, []string{"Py", "GO", "'"})

	for _, term := range terms {
		var lang = term[0]
		var prefix = term[1]
		var quote = term[2]

		fmt.Printf("// %s\n", lang)
		fmt.Println("")
		fmt.Println("// AES Encrypted Values:")
		aes, _ := democrypt.NewAESCrypt(aes128Hash, democrypt.AesTypeCfb)
		encrypted, _ := aes.Encrypt(data)
		fmt.Printf("%s_AES_CFB_128 = %s%s%s\n", prefix, quote, encrypted, quote)
		aes, _ = democrypt.NewAESCrypt(aes192Hash, democrypt.AesTypeCfb)
		encrypted, _ = aes.Encrypt(data)
		fmt.Printf("%s_AES_CFB_192 = %s%s%s\n", prefix, quote, encrypted, quote)
		aes, _ = democrypt.NewAESCrypt(aes256Hash, democrypt.AesTypeCfb)
		encrypted, _ = aes.Encrypt(data)
		fmt.Printf("%s_AES_CFB_256 = %s%s%s\n", prefix, quote, encrypted, quote)
		aes, _ = democrypt.NewAESCrypt(aes128Hash, democrypt.AesTypeCbc)
		encrypted, _ = aes.Encrypt(data)
		fmt.Printf("%s_AES_CBC_128 = %s%s%s\n", prefix, quote, encrypted, quote)
		aes, _ = democrypt.NewAESCrypt(aes192Hash, democrypt.AesTypeCbc)
		encrypted, _ = aes.Encrypt(data)
		fmt.Printf("%s_AES_CBC_192 = %s%s%s\n", prefix, quote, encrypted, quote)
		aes, _ = democrypt.NewAESCrypt(aes256Hash, democrypt.AesTypeCbc)
		encrypted, _ = aes.Encrypt(data)

		fmt.Printf("%s_AES_CBC_256 = %s%s%s\n", prefix, quote, encrypted, quote)
		fmt.Println("")
		cwd, _ := os.Getwd()
		rsaPath := filepath.Join(cwd, "cert", "rsa")

		fmt.Println("// RSA Encrypted Values:")
		rsa, _ := democrypt.NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"), democrypt.RsaOaep)
		encrypted, _ = rsa.Encrypt(data)
		fmt.Printf("%s_RSA_OAEP = %s%s%s\n", prefix, quote, encrypted, quote)
		rsa, _ = democrypt.NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"), democrypt.RsaPkcs1V15)
		encrypted, _ = rsa.Encrypt(data)
		fmt.Printf("%s_RSA_PKCS1V1_5 = %s%s%s\n", prefix, quote, encrypted, quote)

		fmt.Println("")
		x509Path := filepath.Join(cwd, "cert", "x509")

		fmt.Println("// X509 Encrypted Values:")
		x509, _ := democrypt.NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"), democrypt.RsaOaep)
		encrypted, _ = x509.Encrypt(data)
		fmt.Printf("%s_X509_OAEP = %s%s%s\n", prefix, quote, encrypted, quote)
		x509, _ = democrypt.NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"), democrypt.RsaPkcs1V15)
		encrypted, _ = x509.Encrypt(data)
		fmt.Printf("%s_X509_PKCS1V1_5 = %s%s%s\n", prefix, quote, encrypted, quote)

		fmt.Println("")
		fmt.Println("")
	}
}
