package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func CalculateMerkleRoot(txs []*Transaction) string {
	var hashes [][]byte
	for _, tx := range txs {
		txBytes := tx.Hash()
		hashes = append(hashes, txBytes)
	}

	for len(hashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			left := hashes[i]
			var right []byte
			if i+1 < len(hashes) {
				right = hashes[i+1]
			} else {
				right = left // copy nếu số lẻ
			}
			combined := append(left, right...)
			newHash := hashData(combined)
			newLevel = append(newLevel, newHash)
		}
		hashes = newLevel
	}

	if len(hashes) == 0 {
		return ""
	}
	return hex.EncodeToString(hashes[0])
}
