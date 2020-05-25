package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/templ-project/go/src/pkg/democrypt"
)

func doAes(password string) {
	fmt.Println("Using AES:")

	encTool, err := democrypt.NewAESCrypt("1234567890123456")
	if err != nil {
		panic(fmt.Errorf("could not create aes encrypt tool: %v", err.Error()))
	}

	encPasswordAes, _ := encTool.Encrypt(password)
	decPasswordAes, _ := encTool.Decrypt(encPasswordAes)
	// decPasswordAesCsharp, _ := encTool.Decrypt("bJj9fLL39dGw+E926oV23cWB7nIdCF73ojKgfyYEpkg=")

	fmt.Printf("password: %s \nenc password: %s \ndec password: %s \n\n", password, encPasswordAes, decPasswordAes) //, decPasswordAesCsharp)
}

func doSsl(password string) {
	fmt.Println("Using SSL:")

	cwd, _ := os.Getwd()
	sslPath := filepath.Join(cwd, "cert", "ssl")

	sslTool, err := democrypt.NewSSLCrypt(filepath.Join(sslPath, "cert.pem"), filepath.Join(sslPath, "key.pem"))
	if err != nil {
		panic(fmt.Errorf("could not create ssl encrypt tool: %v", err.Error()))
	}

	encPasswordSsl, _ := sslTool.Encrypt(password)
	decPasswordSsl, _ := sslTool.Decrypt(encPasswordSsl)

	// decPasswordSslCSharp, _ := sslTool.Decrypt("j6vg8/BhBjAhruPA2+ajF2+y5H3uyxWAXB3UhjpZ13n42Ba3HVDcTrWatSJbntojKpPtp7AnJ9qgUI3BY9T7en62TC91CkpQeMkSfJPAi1+9jwUgJbGj2jAk5/iO//+Dj46MTCYMH54L1IcRnKobbShEyqtdrcAvUap66BUd0uneKyWktCRrJD8z7T/Yw7klGSal+r/piBwzd3Wl1BIdf3pDVpeVCpEeBBgF+rElit4Id7blr9K+dEtSU70KggXUvCgVWA4FW2KHS6Abkvabv/+9ntT6bH/cFlDBHqo4JJDODPgYtshv82vaEbpSnhbm11VseklUDTUVJDlGf08Gmg==")

	fmt.Printf("password: %s \nenc password: %s \ndec password: %s \n\n", password, encPasswordSsl, decPasswordSsl) //, decPasswordSslCSharp)
}

func doRsa(password string) {
	fmt.Println("Using RSA:")

	cwd, _ := os.Getwd()
	sslPath := filepath.Join(cwd, "cert", "rsa")

	sslTool, err := democrypt.NewRSACrypt(filepath.Join(sslPath, "cert.pem"), filepath.Join(sslPath, "key.pem"))
	if err != nil {
		panic(fmt.Errorf("could not create rsa encrypt tool: %v", err.Error()))
	}

	encPasswordSsl, _ := sslTool.Encrypt(password)
	decPasswordSsl, _ := sslTool.Decrypt(encPasswordSsl)

	// decPasswordSslCSharp, _ := sslTool.Decrypt("j6vg8/BhBjAhruPA2+ajF2+y5H3uyxWAXB3UhjpZ13n42Ba3HVDcTrWatSJbntojKpPtp7AnJ9qgUI3BY9T7en62TC91CkpQeMkSfJPAi1+9jwUgJbGj2jAk5/iO//+Dj46MTCYMH54L1IcRnKobbShEyqtdrcAvUap66BUd0uneKyWktCRrJD8z7T/Yw7klGSal+r/piBwzd3Wl1BIdf3pDVpeVCpEeBBgF+rElit4Id7blr9K+dEtSU70KggXUvCgVWA4FW2KHS6Abkvabv/+9ntT6bH/cFlDBHqo4JJDODPgYtshv82vaEbpSnhbm11VseklUDTUVJDlGf08Gmg==")

	fmt.Printf("password: %s \nenc password: %s \ndec password: %s \n\n", password, encPasswordSsl, decPasswordSsl) //, decPasswordSslCSharp)
}

func main() {
	password := "th1s1smyp@ssw0rd"

	doAes(password)

	doRsa(password)

	doSsl(password)
}
