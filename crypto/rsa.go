package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type Keys struct {
	public  *rsa.PublicKey
	private *rsa.PrivateKey
}

func GenerateRSAKeys(bits int) (*Keys, error) {
	private_key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	return &Keys{
		public:  &private_key.PublicKey,
		private: private_key,
	}, nil
}

func RSAEncrypt(secretmessage []byte, public_key *rsa.PublicKey) ([]byte, error) {
	encrypted_bytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		public_key,
		secretmessage,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return encrypted_bytes, nil
}

func RSADecrypt(cipher_text []byte, private *rsa.PrivateKey) (string, error) {
	decrypted_bytes, err := private.Decrypt(
		rand.Reader,
		cipher_text,
		&rsa.OAEPOptions{Hash: crypto.SHA256},
	)
	if err != nil {
		panic(err)
	}
	return string(decrypted_bytes), nil
}
