package twowaykey

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func PemPrivDecode(pemPrivKey string) ([]byte, error) {

	block, _ := pem.Decode([]byte(pemPrivKey))

	if block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("decoded pem not of type rsa private key")
	}

	return block.Bytes, nil
}

func PemPubDecode(pemPubKey string) ([]byte, error) {

	block, _ := pem.Decode([]byte(pemPubKey))

	if block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("decoded pem not of type rsa public key")
	}

	return block.Bytes, nil
}

func ParsePrivKey(bytePrivateKey []byte) (*rsa.PrivateKey, error) {

	privateKey, err := x509.ParsePKCS1PrivateKey(bytePrivateKey)

	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa private key : %w", err)
	}

	return privateKey, nil
}

func ParsePubKey(bytePublicKey []byte) (*rsa.PublicKey, error) {

	publicKeyIface, err := x509.ParsePKIXPublicKey(bytePublicKey)

	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa public key : %w", err)
	}

	publicKey, ok := publicKeyIface.(*rsa.PublicKey)

	if !ok {
		return nil, fmt.Errorf("interface does not contain rsa public key")
	}

	return publicKey, nil
}

func DecAndParsePrivKey(pemPrivKey string) (*rsa.PrivateKey, error) {

	bytePrivateKey, err := PemPrivDecode(pemPrivKey)

	if err != nil {
		return nil, err
	}

	privateKey, err := ParsePrivKey(bytePrivateKey)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func DecAndParsePubKey(pemPubKey string) (*rsa.PublicKey, error) {

	bytePublicKey, err := PemPubDecode(pemPubKey)

	if err != nil {
		return nil, err
	}

	publicKey, err := ParsePubKey(bytePublicKey)

	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
