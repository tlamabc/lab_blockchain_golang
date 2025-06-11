package main

import (
	"fmt"
	"os"
	"time"
	"encoding/json"

	"interview-be-earning/pkg/blockchain"
	"interview-be-earning/pkg/wallet"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli create|load <file>")
		return
	}

	command := os.Args[1]
	filename := "wallet.json"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}

	switch command {
	case "create":
		w, err := wallet.NewWallet()
		if err != nil {
			panic(err)
		}
		if err := wallet.SaveWallet(w, filename); err != nil {
			panic(err)
		}
		fmt.Println("Wallet created and saved to", filename)
		fmt.Println("Address:", wallet.PublicKeyToAddress(w.PublicKey))
	case "load":
		w, err := wallet.LoadWallet(filename)
		if err != nil {
			panic(err)
		}
		fmt.Println("Wallet loaded from", filename)
		fmt.Println("Address:", wallet.PublicKeyToAddress(w.PublicKey))
	default:
		fmt.Println("Unknown command")

	///
	case "sendtx":
	w, err := wallet.LoadWallet(filename)
	if err != nil {
		panic(err)
	}
	tx := &blockchain.Transaction{
		Sender:    wallet.PublicKeyToAddress(w.PublicKey),
		Receiver:  "bob-address-xyz",
		Amount:    10.5,
		Timestamp: time.Now().Unix(),
	}
	if err := tx.Sign(w.PrivateKey); err != nil {
		panic(err)
	}
	fmt.Println("Signed TX:")
	txJSON, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println(string(txJSON))

	// Thử verify luôn
	ok, err := tx.Verify(&w.PublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Verify status:", ok)
	///


	
	case "makeblock":
	w, _ := wallet.LoadWallet(filename)

	// Tạo 2 giao dịch
	tx1 := &blockchain.Transaction{
		Sender:    wallet.PublicKeyToAddress(w.PublicKey),
		Receiver:  "bob-address-xyz",
		Amount:    5,
		Timestamp: time.Now().Unix(),
	}
	tx1.Sign(w.PrivateKey)

	tx2 := &blockchain.Transaction{
		Sender:    wallet.PublicKeyToAddress(w.PublicKey),
		Receiver:  "bob-address-abc",
		Amount:    15,
		Timestamp: time.Now().Unix(),
	}
	tx2.Sign(w.PrivateKey)

	// Tạo block từ 2 tx
	block := blockchain.NewBlock([]*blockchain.Transaction{tx1, tx2}, "prev_hash_dummy")

	// In block
	blockJson, _ := json.MarshalIndent(block, "", "  ")
	fmt.Println(string(blockJson))

	}



		node := &Node{
		ID:   os.Getenv("NODE_ID"),
		Role: os.Getenv("ROLE"),
		Port: os.Getenv("PORT"),
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from %s", node.ID)
	})

	http.HandleFunc("/submit-tx", node.handleSubmitTx)

	log.Println("✅ Starting node:", node.ID, "on port", node.Port)
	log.Fatal(http.ListenAndServe(":"+node.Port, nil))
}
