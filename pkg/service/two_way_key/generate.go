package twowaykey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate rsa private and public key : %w", err)
	}

	return privateKey, &privateKey.PublicKey, nil
}

func MarshalPrivateKey(privateKey *rsa.PrivateKey) ([]byte, error) {

	var bytePrivateKey = x509.MarshalPKCS1PrivateKey(privateKey)

	if bytePrivateKey == nil {
		return nil, fmt.Errorf("failed to marshal rsa private key")
	}

	return bytePrivateKey, nil
}

func MarshalPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {

	bytePublicKey, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal rsa public key : %w", err)
	}

	return bytePublicKey, nil
}

func PemPrivEncode(bytePrivateKey []byte) (string, error) {

	var pemEncode = pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: bytePrivateKey,
	})

	if pemEncode == nil {
		return "", fmt.Errorf("failed to pem encode rsa private key")
	}

	return string(pemEncode), nil
}

func PemPubEncode(bytePublicKey []byte) (string, error) {

	var pemEncode = pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY", Bytes: bytePublicKey,
	})

	if pemEncode == nil {
		return "", fmt.Errorf("failed to pem encode rsa public key")
	}

	return string(pemEncode), nil
}

func GenAndEncode() (string, string, error) {

	privateKey, publicKey, err := GenerateKeys()

	if err != nil {
		return "", "", err
	}

	bytePrivateKey, err := MarshalPrivateKey(privateKey)

	if err != nil {
		return "", "", err
	}

	bytePublicKey, err := MarshalPublicKey(publicKey)

	if err != nil {
		return "", "", err
	}

	strPemPrivKey, err := PemPrivEncode(bytePrivateKey)

	if err != nil {
		return "", "", err
	}

	strPemPubKey, err := PemPubEncode(bytePublicKey)

	if err != nil {
		return "", "", err
	}

	return strPemPrivKey, strPemPubKey, nil
}
