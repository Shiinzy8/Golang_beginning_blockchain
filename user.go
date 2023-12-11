package Golang_beginning_blockchain

import "errors"

type User struct {
	name      string
	id        int
	amountDol float64
	amountBit float64
}

func (u User) checkAmountDol(amountDol float64) (bool, error) {
	if u.amountDol > amountDol {
		return true, nil
	} else {
		err := errors.New("math: Not enough money on the account")
		return false, err
	}
}

func (u User) checkAmountBit(amountBit float64) (bool, error) {
	if u.amountBit > amountBit {
		return true, nil
	} else {
		err := errors.New("math: Not enough bitcoins on the account")
		return false, err
	}
}
