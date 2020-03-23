package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)
import "time"

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

type Block struct {
	Index        int64
	Timestamp    time.Time
	Transactions []Transaction
	Proof        string
	PreviousHash string
}

type Blockchain struct {
	chain               []Block
	currentTransactions []Transaction
}

func (bc *Blockchain) newBlock(proof string, previousHash string) {

	if previousHash == "" {
		previousHash = hash(bc.chain[0])
	}

	block := Block{
		Index:        int64(len(bc.chain)) + 1,
		Timestamp:    time.Now(),
		Transactions: bc.currentTransactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}

	bc.currentTransactions = nil

	bc.chain = append(bc.chain, block)
}

func (bc *Blockchain) newTransaction(sender string, recipient string, amount float64) int {
	trx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	bc.currentTransactions = append(bc.currentTransactions, trx)

	return len(bc.currentTransactions)
}

func hash(block Block) string {
	blockString, _ := json.Marshal(block)

	h := sha256.New()
	h.Write(blockString)
	hash := h.Sum(nil)
	str := hex.EncodeToString(hash)
	return str
}

func (bc *Blockchain) proofOfWork(lastProof string) int {
	proof := 0

	for !validProof(lastProof, string(proof)) {
		proof += 1
	}

	return proof
}

func validProof(lastProof string, proof string) bool {
	guess := lastProof + proof
	h := sha256.New()
	h.Write([]byte(guess))
	hash := h.Sum(nil)
	str := hex.EncodeToString(hash)
	if str[0:5] == "00000" {
		return true
	}
	return false
}

func main() {
	bc := Blockchain{
		chain:               nil,
		currentTransactions: nil,
	}

	bc.newTransaction("mehran", "test", 1000)
	bc.newTransaction("z", "x", 2000)
	test := hash(bc.chain[0])
	fmt.Println(test)
	//
	//bc.newBlock()
	proof := bc.proofOfWork("dsadsadwquye7ywqdbsabdhsa")
	fmt.Println(proof)
}
