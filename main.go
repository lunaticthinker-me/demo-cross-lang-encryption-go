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

func formatError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%v", err)
}

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

func getAes(key string, _type int) (*democrypt.AESCrypt, error) {
	return democrypt.NewAESCrypt(key, _type)
}

func indexOf(element string, data map[string]int) int {
	for k, v := range data {
		if element == k {
			return v
		}
	}
	return -1
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
	for cypherLabel, cypherType := range democrypt.AesChyperModes {
		keys := []string{}
		k, _ := randutil.Alphanumeric(16)
		keys = append(keys, k)
		k, _ = randutil.Alphanumeric(24)
		keys = append(keys, k)
		k, _ = randutil.Alphanumeric(32)
		keys = append(keys, k)
		for _, key := range keys {
			var row []string
			algo := fmt.Sprintf("%s:%s", "AES", cypherLabel)
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
						fmt.Sprintf("%v", formatError(err)),
					}
				} else {
					row = []string{algo, key, decrypted, "", fmt.Sprintf("%v", formatError(err))}
				}
			} else {
				row = []string{algo, key, "", "", fmt.Sprintf("%v", formatError(err))}
			}

			tbl.AddRow(row[0], row[1], row[2], row[3], row[4])
			csv = append(csv, row)
		}
	}

	// RSA

	for paddingLabel, padding := range democrypt.RsaPaddings {
		var row []string
		algo := fmt.Sprintf("%s:%s", "RSA", paddingLabel)
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
					fmt.Sprintf("%v", formatError(err)),
				}
			} else {
				row = []string{algo, "", decrypted, "", fmt.Sprintf("%v", formatError(err))}
			}
		} else {
			row = []string{algo, "", "", "", fmt.Sprintf("%v", formatError(err))}
		}

		tbl.AddRow(row[0], row[1], row[2], row[3], row[4])
		csv = append(csv, row)
	}

	// X509

	for paddingLabel, padding := range democrypt.RsaPaddings {
		var row []string
		algo := fmt.Sprintf("%s:%s", "X509", paddingLabel)
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
					fmt.Sprintf("%v", formatError(err)),
				}
			} else {
				row = []string{algo, "", decrypted, "", fmt.Sprintf("%v", formatError(err))}
			}
		} else {
			row = []string{algo, "", "", "", fmt.Sprintf("%v", formatError(err))}
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

	csvString := strings.Trim(string(csvBytes), "\n")

	for _, line := range strings.Split(csvString, "\n") {
		test := strings.Split(line, ",")
		algo := strings.Split(test[0], ":")
		var crypto democrypt.Crypt
		var err error

		switch algo[0] {
		case "AES":
			crypto, err = getAes(test[1], indexOf(algo[1], democrypt.AesChyperModes))
		case "RSA":
			crypto, err = getRsa(indexOf(algo[1], democrypt.RsaPaddings))
		case "X509":
			crypto, err = getX509(indexOf(algo[1], democrypt.RsaPaddings))
		}

		if test[4] == "" {
			if err == nil {
				decrypted, err := crypto.Decrypt(test[3])
				if err == nil {
					if decrypted == test[2] {
						tbl.AddRow(test[0], "yes", "")
					} else {
						tbl.AddRow(test[0], "no", "decrypt failed")
					}
				} else {
					tbl.AddRow(test[0], "no", fmt.Sprintf("%v", formatError(err)))
				}
			} else {
				tbl.AddRow(test[0], "no", fmt.Sprintf("%v", formatError(err)))
			}
		} else {
			tbl.AddRow(test[0], "no", fmt.Sprintf("source could not encrypt: %s", test[4]))
		}
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
