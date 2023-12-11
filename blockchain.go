package Golang_beginning_blockchain

import (
	"time"
)

type Operation string

const (
	AddDol        Operation = "Add dollars to account"
	ConvertDolBit Operation = "Convert dollars to bitcoins"
	ConvertBitDol Operation = "Convert bitcoins to dollars"
	RemoveDol     Operation = "Remove dollars from account"
	SentBit       Operation = "Sent bitcoins to another user"
)

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
	genesisUser  User
	users        []User
	countUsers   int
}

func (b *Blockchain) AddDol(user User, amountDol float64) Block {
	user.addDol(amountDol)
	return b.addBlock(b.genesisUser, user, AddDol, amountDol, 0)
}

func (b *Blockchain) RemoveDol(user User, amountDol float64) (Block, error) {
	if ok, err := user.checkAmountDol(amountDol); ok {
		user.removeDol(amountDol)
		return b.addBlock(user, b.genesisUser, RemoveDol, amountDol, 0), nil
	} else {
		return Block{}, err
	}
}

func (b *Blockchain) SentBit(fromUser, toUser User, amountBit float64) (Block, error) {
	if ok, err := fromUser.checkAmountBit(amountBit); ok {
		fromUser.removeBit(amountBit)
		toUser.addBit(amountBit)
		return b.addBlock(fromUser, toUser, SentBit, 0, amountBit), nil
	} else {
		return Block{}, err
	}
}

func (b *Blockchain) ConvertDolBit(user User, amountDol float64) (Block, error) {
	if ok, err := user.checkAmountDol(amountDol); ok {
		user.removeDol(amountDol)
		amountBit := amountDol / 36000
		user.addBit(amountBit)
		return b.addBlock(user, user, ConvertDolBit, amountDol, amountBit), nil
	} else {
		return Block{}, err
	}
}

func (b *Blockchain) ConvertBitDol(user User, amountBit float64) (Block, error) {
	if ok, err := user.checkAmountBit(amountBit); ok {
		user.removeBit(amountBit)
		amountDol := amountBit * 36000
		user.addDol(amountDol)
		return b.addBlock(user, user, ConvertBitDol, amountDol, amountBit), nil
	} else {
		return Block{}, err
	}
}

func (b *Blockchain) addBlock(from, to User, operation Operation, amountDol float64, amountBit float64) Block {
	blockData := BlockData{
		from:      from,
		to:        to,
		operation: operation,
		amountDol: amountDol,
		amountBit: amountBit,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.Mine(b.difficulty)
	b.chain = append(b.chain, newBlock)

	return newBlock
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	genesisUser := User{
		name: "Blockchain",
		id:   0,
	}

	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
		genesisUser,
		[]User{genesisUser},
		0,
	}
}

func (b *Blockchain) AddUser(name string) User {
	b.countUsers++
	newUser := User{name: name, id: b.countUsers, amountBit: 0, amountDol: 0}
	b.users = append(b.users, newUser)

	return newUser
}

func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (b Blockchain) GetChain() []Block {
	return b.chain
}

func (b Blockchain) GetGenesisBlock() Block {
	return b.genesisBlock
}

func (b Blockchain) GetUsers() []User {
	return b.users
}

func (b Blockchain) GetGenesisUser() User {
	return b.genesisUser
}

func (b Blockchain) GetDifficulty() int {
	return b.difficulty
}
