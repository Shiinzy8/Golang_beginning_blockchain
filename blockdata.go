package Golang_beginning_blockchain

import "encoding/json"

type BlockData struct {
	from      User
	to        User
	amountDol float64
	amountBit float64
	operation Operation
	json.Marshaler
}

func (b BlockData) toJson() ([]byte, error) {
	data := map[string]interface{}{
		"from":      b.from.id,
		"to":        b.to.id,
		"operation": b.operation,
		"amountDol": b.amountDol,
		"amountBit": b.amountBit,
	}

	return json.Marshal(data)
}

func (b BlockData) GetFromUser() User {
	return b.from
}

func (b BlockData) GetToUser() User {
	return b.to
}

func (b BlockData) GetAmountDol() float64 {
	return b.amountDol
}

func (b BlockData) GetAmountBit() float64 {
	return b.amountDol
}
