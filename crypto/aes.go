package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AESKey []byte

func GenerateKey(bits int) (AESKey, error) {
	key := make([]byte, bits)

	_, err := io.ReadFull(rand.Reader, key)

	if err != nil {
		panic(err.Error())
	}

	return key, nil
}

func AESEncrypt(secret_message []byte, key AESKey) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// defining Galois Counting Mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//defining nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipher_text := aesgcm.Seal(nil, nonce, secret_message, nil)
	return cipher_text, nil
}

func AESDecrypt(cipher_text []byte, key AESKey) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	decrypted_bytes, err := aesgcm.Open(nil, cipher_text[:aesgcm.NonceSize()], cipher_text, nil)
	if err != nil {
		panic(err.Error())
	}

	return decrypted_bytes, nil
}
