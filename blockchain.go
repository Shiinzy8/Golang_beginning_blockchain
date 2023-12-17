package Golang_beginning_blockchain

import (
	"fmt"
	"time"
)

type operation string

const (
	addDol        operation = "Add dollars to account"
	convertDolBit operation = "Convert dollars to bitcoins"
	convertBitDol operation = "Convert bitcoins to dollars"
	removeDol     operation = "Remove dollars from account"
	sentBit       operation = "Sent bitcoins to another user"
)

type Blockchain struct {
	genesisBlock block
	chain        []block
	difficulty   int
	genesisUser  *user
	users        map[int]*user
	countUsers   int
}

func (b *Blockchain) AddDol(user *user, amountDol float64) block {
	b.users[user.id].addDol(amountDol)
	return b.addBlock(b.genesisUser, user, addDol, amountDol, 0)
}

func (b *Blockchain) RemoveDol(user *user, amountDol float64) (block, error) {
	if ok, err := b.users[user.id].checkAmountDol(amountDol); ok {
		b.users[user.id].removeDol(amountDol)
		return b.addBlock(user, b.genesisUser, removeDol, amountDol, 0), nil
	} else {
		return block{}, err
	}
}

func (b *Blockchain) SentBit(fromUser, toUser *user, amountBit float64) (block, error) {
	if ok, err := b.users[fromUser.id].checkAmountBit(amountBit); ok {
		b.users[fromUser.id].removeBit(amountBit)
		b.users[toUser.id].addBit(amountBit)
		return b.addBlock(fromUser, toUser, sentBit, 0, amountBit), nil
	} else {
		return block{}, err
	}
}

func (b *Blockchain) ConvertDolBit(user *user, amountDol float64) (block, error) {
	if ok, err := b.users[user.id].checkAmountDol(amountDol); ok {
		b.users[user.id].removeDol(amountDol)
		amountBit := amountDol / 36000
		b.users[user.id].addBit(amountBit)
		return b.addBlock(user, user, convertDolBit, amountDol, amountBit), nil
	} else {
		return block{}, err
	}
}

func (b *Blockchain) ConvertBitDol(user *user, amountBit float64) (block, error) {
	if ok, err := b.users[user.id].checkAmountBit(amountBit); ok {
		b.users[user.id].removeBit(amountBit)
		amountDol := amountBit * 36000
		b.users[user.id].addDol(amountDol)
		return b.addBlock(user, user, convertBitDol, amountDol, amountBit), nil
	} else {
		return block{}, err
	}
}

func (b *Blockchain) addBlock(from, to *user, operation operation, amountDol float64, amountBit float64) block {
	blockData := blockData{
		from:      *from,
		to:        *to,
		operation: operation,
		amountDol: amountDol,
		amountBit: amountBit,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.Mine(b.difficulty)
	b.chain = append(b.chain, newBlock)

	return newBlock
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := block{
		hash:      "0",
		timestamp: time.Now(),
	}

	genesisUser := &user{
		name: "Blockchain",
		id:   0,
	}

	return Blockchain{
		genesisBlock,
		[]block{genesisBlock},
		difficulty,
		genesisUser,
		map[int]*user{0: genesisUser},
		0,
	}
}

func (b *Blockchain) AddUser(name string) *user {
	b.countUsers++
	newUser := user{name: name, id: b.countUsers, amountBit: 0, amountDol: 0}
	b.users[b.countUsers] = &newUser

	return &newUser
}

func (b Blockchain) IsValid() bool {
	for i := 0; i < len(b.chain)-1; i++ {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (b Blockchain) GetChain() []block {
	return b.chain
}

func (b Blockchain) GetGenesisBlock() block {
	return b.genesisBlock
}

func (b Blockchain) GetUsers() map[int]*user {
	return b.users
}

func (b Blockchain) GetGenesisUser() *user {
	return b.genesisUser
}

func (b Blockchain) GetDifficulty() int {
	return b.difficulty
}

func (b Blockchain) PrintUsers() {
	for _, user := range b.users {
		fmt.Println(user.name)
		fmt.Println(user.amountBit, "Bit")
		fmt.Println(user.amountDol, "Dol")
	}
}
