package Golang_beginning_blockchain

import "errors"

type user struct {
	name      string
	id        int
	amountDol float64
	amountBit float64
}

func (u user) checkAmountDol(amountDol float64) (bool, error) {
	if u.amountDol > amountDol {
		return true, nil
	} else {
		err := errors.New("math: Not enough money on the account")
		return false, err
	}
}

func (u user) checkAmountBit(amountBit float64) (bool, error) {
	if u.amountBit > amountBit {
		return true, nil
	} else {
		err := errors.New("math: Not enough bitcoins on the account")
		return false, err
	}
}

func (u *user) addDol(amountDol float64) {
	u.amountDol += amountDol
}

func (u *user) removeDol(amountDol float64) {
	u.amountDol -= amountDol
}

func (u *user) addBit(amountBit float64) {
	u.amountBit += amountBit
}

func (u *user) removeBit(amountBit float64) {
	u.amountBit -= amountBit
}

func (u user) GetAmountDol() float64 {
	return u.amountDol
}

func (u user) GetAmountBit() float64 {
	return u.amountBit
}
