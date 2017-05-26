package vapidkeys

import (
	"crypto/elliptic"
	"encoding/base64"
	"github.com/enceve/crypto/dh/ecdh"
)

func Generate() (string, string, error) {
	curve := ecdh.GenericCurve(elliptic.P256())
	privateKey, publicKey, err := curve.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	vapidPrivateKey := base64.RawURLEncoding.EncodeToString(privateKey)
	vapidPublicKey := base64.RawURLEncoding.EncodeToString(publicKey)
	return vapidPrivateKey, vapidPublicKey, nil
}
