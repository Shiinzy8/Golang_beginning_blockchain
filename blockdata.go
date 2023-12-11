package Golang_beginning_blockchain

import "encoding/json"

type BlockData struct {
	from   string
	to     string
	amount float64
	json.Marshaler
}

func (b BlockData) toJson() ([]byte, error) {
	data := map[string]interface{}{
		"from":   b.from,
		"to":     b.to,
		"amount": b.amount,
	}

	return json.Marshal(data)
}

func (b BlockData) GetFrom() string {
	return b.from
}

func (b BlockData) GetTo() string {
	return b.to
}

func (b BlockData) GetAmount() float64 {
	return b.amount
}
