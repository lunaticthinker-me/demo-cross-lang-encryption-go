package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	democrypt "github.com/lunaticthinker-me/demo-cross-lang-encryption-go/src"
	"github.com/rodaine/table"
	"go.step.sm/crypto/randutil"
)

func getRsa(Type int) (*democrypt.RsaCrypt, error) {
	cwd, _ := os.Getwd()
	rsaPath := filepath.Join(cwd, "cert", "rsa")
	return democrypt.NewRSACrypt(filepath.Join(rsaPath, "cert.pem"), filepath.Join(rsaPath, "key.pem"), Type)
}

func getX509(Type int) (*democrypt.X509Crypt, error) {
	cwd, _ := os.Getwd()
	x509Path := filepath.Join(cwd, "cert", "x509")
	return democrypt.NewX509Crypt(filepath.Join(x509Path, "cert.pem"), filepath.Join(x509Path, "key.pem"), Type)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func writeCsv(csv string) {
	f, err := os.Create("data.csv")
	panicIfErr(err)
	defer f.Close()

	_, err = f.WriteString(csv)
	panicIfErr(err)
}

func Encrypt() {
	tbl := table.New("Algo", "Key", "Decrypted", "Encrypted", "EncryptionError")
	csv := [][]string{}
	csvString := ""

	// AES
	for _, cypherType := range democrypt.AesChyperList {
		keys := []string{}
		k, _ := randutil.Alphanumeric(16)
		keys = append(keys, k)
		k, _ = randutil.Alphanumeric(24)
		keys = append(keys, k)
		k, _ = randutil.Alphanumeric(32)
		keys = append(keys, k)
		for _, key := range keys {
			var row []string
			algo := fmt.Sprintf("%s:%s", "AES", democrypt.AesCypherLabels[cypherType])
			decrypted, err := randutil.Alphanumeric(16)
			if err == nil {
				crypto, err := democrypt.NewAESCrypt(key, cypherType)
				if err == nil {
					encrypted, err := crypto.Encrypt(decrypted)

					row = []string{
						algo,
						key,
						decrypted,
						encrypted,
						fmt.Sprintf("%v", err),
					}
				} else {
					row = []string{algo, key, decrypted, "", fmt.Sprintf("%v", err)}
				}
			} else {
				row = []string{algo, key, "", "", fmt.Sprintf("%v", err)}
			}

			tbl.AddRow(row[0], row[1], row[2], row[3], row[4])
			csv = append(csv, row)
		}
	}

	// RSA

	for _, padding := range democrypt.RsaPaddingList {
		var row []string
		algo := fmt.Sprintf("%s:%s", "RSA", democrypt.RsaPaddingLabels[padding])
		decrypted, err := randutil.Alphanumeric(16)
		if err == nil {
			crypto, err := getRsa(padding)
			if err == nil {
				encrypted, err := crypto.Encrypt(decrypted)

				row = []string{
					algo,
					"",
					decrypted,
					encrypted,
					fmt.Sprintf("%v", err),
				}
			} else {
				row = []string{algo, "", decrypted, "", fmt.Sprintf("%v", err)}
			}
		} else {
			row = []string{algo, "", "", "", fmt.Sprintf("%v", err)}
		}

		tbl.AddRow(row[0], row[1], row[2], row[3], row[4])
		csv = append(csv, row)
	}

	// X509

	for _, padding := range democrypt.RsaPaddingList {
		var row []string
		algo := fmt.Sprintf("%s:%s", "X509", democrypt.RsaPaddingLabels[padding])
		decrypted, err := randutil.Alphanumeric(16)
		if err == nil {
			crypto, err := getX509(padding)
			if err == nil {
				encrypted, err := crypto.Encrypt(decrypted)

				row = []string{
					algo,
					"",
					decrypted,
					encrypted,
					fmt.Sprintf("%v", err),
				}
			} else {
				row = []string{algo, "", decrypted, "", fmt.Sprintf("%v", err)}
			}
		} else {
			row = []string{algo, "", "", "", fmt.Sprintf("%v", err)}
		}

		tbl.AddRow(row[0], row[1], row[2], row[3], row[4])
		csv = append(csv, row)
	}

	tbl.Print()

	for _, row := range csv {
		csvString = csvString + strings.Join(row, ",") + "\n"
	}

	writeCsv(strings.Trim(csvString, "\n"))
}

func Decrypt() {
	csvBytes, err := ioutil.ReadAll(os.Stdin)
	panicIfErr(err)

	tbl := table.New("Algo", "Decrypted", "DecryptionError")

	csvString := string(csvBytes)

	for _, line := range strings.Split(csvString, "\n") {
		test := strings.Split(line, ",")
		// algo := strings.Split(test[0], ":")

		tbl.AddRow(test, "", "")
	}

	tbl.Print()
}

func main() {
	command := os.Args[1]

	switch command {
	case "decrypt":
		Decrypt()
	case "encrypt":
		Encrypt()
	default:
		fmt.Println("invalid command")
	}
}