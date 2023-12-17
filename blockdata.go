package Golang_beginning_blockchain

import "encoding/json"

type blockData struct {
	from      user
	to        user
	amountDol float64
	amountBit float64
	operation operation
	json.Marshaler
}

func (b blockData) toJson() ([]byte, error) {
	data := map[string]interface{}{
		"from":      b.from.id,
		"to":        b.to.id,
		"operation": b.operation,
		"amountDol": b.amountDol,
		"amountBit": b.amountBit,
	}

	return json.Marshal(data)
}

func (b blockData) GetFromUser() user {
	return b.from
}

func (b blockData) GetToUser() user {
	return b.to
}

func (b blockData) GetAmountDol() float64 {
	return b.amountDol
}

func (b blockData) GetAmountBit() float64 {
	return b.amountDol
}
