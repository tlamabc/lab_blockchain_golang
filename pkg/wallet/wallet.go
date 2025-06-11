package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"os"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

func NewWallet() (*Wallet, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return &Wallet{PrivateKey: priv, PublicKey: priv.PublicKey}, nil
}

func SaveWallet(w *Wallet, filename string) error {
	data := map[string]string{
		"D": hex.EncodeToString(w.PrivateKey.D.Bytes()),
		"X": hex.EncodeToString(w.PublicKey.X.Bytes()),
		"Y": hex.EncodeToString(w.PublicKey.Y.Bytes()),
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func LoadWallet(filename string) (*Wallet, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var fields map[string]string
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}

	dBytes, _ := hex.DecodeString(fields["D"])
	xBytes, _ := hex.DecodeString(fields["X"])
	yBytes, _ := hex.DecodeString(fields["Y"])

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()
	priv.D = new(big.Int).SetBytes(dBytes)
	priv.PublicKey.X = new(big.Int).SetBytes(xBytes)
	priv.PublicKey.Y = new(big.Int).SetBytes(yBytes)

	return &Wallet{PrivateKey: priv, PublicKey: priv.PublicKey}, nil
}

func PublicKeyToAddress(pub ecdsa.PublicKey) string {
	pubKeyBytes := append(pub.X.Bytes(), pub.Y.Bytes()...)
	hash := sha256.Sum256(pubKeyBytes)
	return hex.EncodeToString(hash[:])
}
