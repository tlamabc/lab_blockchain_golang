package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"math/big"
	"crypto/rand"
)

type Transaction struct {
	Sender    string  `json:"sender"`   // Địa chỉ public key của người gửi (hash của public key)
	Receiver  string  `json:"receiver"` // Địa chỉ người nhận
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
	Signature []byte  `json:"signature"` // R || S nối lại
}
func PublicKeyFromAddress(addr string) (*ecdsa.PublicKey, error) {
	// Tạm thời trả về lỗi — chưa map từ address về public key thật.
	return nil, errors.New("not implemented")
}
// Hash giao dịch (bỏ chữ ký ra)
func (tx *Transaction) Hash() []byte {
	txCopy := *tx
	txCopy.Signature = nil
	data, _ := json.Marshal(txCopy)
	hash := sha256.Sum256(data)
	return hash[:]
}

// Ký giao dịch bằng private key
func (tx *Transaction) Sign(priv *ecdsa.PrivateKey) error {
	hash := tx.Hash()
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)

	// r, s, err := ecdsa.Sign(nil, priv, hash)
	if err != nil {
		return err
	}
	// Nối r và s lại làm signature
	tx.Signature = append(r.Bytes(), s.Bytes()...)
	return nil
}

// Xác thực chữ ký bằng public key
func (tx *Transaction) Verify(pub *ecdsa.PublicKey) (bool, error) {
	if tx.Signature == nil {
		return false, errors.New("no signature")
	}
	hash := tx.Hash()

	// Tách signature → r, s
	sigLen := len(tx.Signature)
	r := new(big.Int).SetBytes(tx.Signature[:sigLen/2])
	s := new(big.Int).SetBytes(tx.Signature[sigLen/2:])

	valid := ecdsa.Verify(pub, hash, r, s)
	return valid, nil
}
