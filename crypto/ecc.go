package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

type EllipticCurve struct {
	pubKeyCurve elliptic.Curve
	privateKey  *ecdsa.PrivateKey
	publicKey   *ecdsa.PublicKey
}

func NewCurve(curve elliptic.Curve) *EllipticCurve {
	return &EllipticCurve{
		pubKeyCurve: curve,
		privateKey:  new(ecdsa.PrivateKey),
	}
}

func (ec *EllipticCurve) GenerateEccKeys() (privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, err error) {
	privateKey, err = ecdsa.GenerateKey(ec.pubKeyCurve, rand.Reader)

	if err != nil {
		panic(err.Error())
	}

	return privateKey, &privateKey.PublicKey, nil

}

func (ec *EllipticCurve) EncodeEccPrivateKey(privKey *ecdsa.PrivateKey) (key string, err error) {
	encoded, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		panic(err.Error())
	}
	// private enhaced mails setup to send the key to the server
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "Private ECC Key", Bytes: encoded})

	key = string(pemEncoded)
	return key, nil
}

func (ec *EllipticCurve) EncodeEccPublicKey(pubKey *ecdsa.PublicKey) (key string, err error) {
	encoded, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err.Error())
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "Public ECC Key", Bytes: encoded})
	key = string(pemEncoded)
	return key, nil
}
