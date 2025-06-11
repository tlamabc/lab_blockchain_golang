package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Transactions      []*Transaction `json:"transactions"`
	MerkleRoot        string         `json:"merkle_root"`
	PreviousBlockHash string         `json:"previous_block_hash"`
	Hash              string         `json:"hash"`
}

// Tính hash toàn bộ block
func (b *Block) CalculateHash() string {
	blockCopy := *b
	blockCopy.Hash = "" // tránh vòng lặp vô hạn
	data, _ := json.Marshal(blockCopy)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// Tạo block mới
func NewBlock(txs []*Transaction, prevHash string) *Block {
	root := CalculateMerkleRoot(txs)
	block := &Block{
		Transactions:      txs,
		MerkleRoot:        root,
		PreviousBlockHash: prevHash,
	}
	block.Hash = block.CalculateHash()
	return block
}
