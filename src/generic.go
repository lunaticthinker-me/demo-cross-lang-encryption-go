package democrypt

type Crypt interface {
	// Encrypt will encrypt a string using Crypto algorithm, returning a Base64 form of the result
	Encrypt(plaintext string) (string, error)

	// EncryptBytes will encrypt a set of bytes using Crypto algorithm, returning a Base64 form of the result
	EncryptBytes(plaintext []byte) ([]byte, error)

	// Decrypt will decrypt a Base64 encoded string using Crypto algorithm returning a string
	Decrypt(ciphertext string) (string, error)

	// DecryptBytes will decrypt a set of Base64 encoded bytes using Crypto algorithm returning a byte array
	DecryptBytes(cipherbytes []byte) ([]byte, error)
}
