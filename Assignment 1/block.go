package main

//  go run main.go block.go
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         CalculateHash(transaction + string(nonce) + previousHash),
	}
	Blockchain = append(Blockchain, *block)
	return block
}

func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

func ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(Blockchain) {
		Blockchain[blockIndex].Transaction = newTransaction
		Blockchain[blockIndex].Hash = CalculateHash(newTransaction + string(Blockchain[blockIndex].Nonce) + Blockchain[blockIndex].PreviousHash)
	} else {
		fmt.Println("Invalid block index.")
	}
}

func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			return false
		}
	}
	return true
}
