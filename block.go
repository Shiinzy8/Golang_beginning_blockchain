package Golang_beginning_blockchain

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type block struct {
	data         blockData
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

func (b block) calculateHash() string {
	json, _ := b.data.toJson()
	blockData := b.previousHash + string(json) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func (b *block) GetHash() string {
	return b.hash
}

func (b *block) GetPreviousHash() string {
	return b.previousHash
}

func (b *block) GetPow() int {
	return b.pow
}

func (b *block) GetCreatedTime() time.Time {
	return b.timestamp
}

func (b *block) GetData() blockData {
	return b.data
}
