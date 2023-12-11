package Golang_beginning_blockchain

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         BlockData
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

func (b Block) CalculateHash() string {
	json, _ := b.data.toJson()
	blockData := b.previousHash + string(json) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

func (b *Block) GetHash() string {
	return b.hash
}

func (b *Block) GetPreviousHash() string {
	return b.previousHash
}

func (b *Block) GetPow() int {
	return b.pow
}

func (b *Block) GetCreatedTime() time.Time {
	return b.timestamp
}

func (b *Block) GetData() BlockData {
	return b.data
}
