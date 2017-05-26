package vapidkeys

import (
	"crypto/elliptic"
	"encoding/base64"
	"github.com/enceve/crypto/dh/ecdh"
)

// This function generates a new set of vapid keys and returns them along with the error message
func Generate() (vapidPrivateKey string, vapidPublicKey string, err error) {
	curve := ecdh.GenericCurve(elliptic.P256())
	privateKey, publicKey, err := curve.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	privKey := base64.RawURLEncoding.EncodeToString(privateKey)
	pubKey := base64.RawURLEncoding.EncodeToString(publicKey)
	return privKey, pubKey, nil
}
