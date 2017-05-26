package vapidkeys

import (
	"encoding/base64"
	"testing"
)

func TestKeys(t *testing.T) {
	privKey, pubKey, err := Generate()
	if err != nil {
		t.Error(err.Error())
	}
	_, err = base64.RawURLEncoding.DecodeString(privKey)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = base64.RawURLEncoding.DecodeString(pubKey)
	if err != nil {
		t.Error(err.Error())
	}
}
