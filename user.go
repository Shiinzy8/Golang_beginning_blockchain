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

func (u *User) addDol(amountDol float64) {
	u.amountDol += amountDol
}

func (u *User) removeDol(amountDol float64) {
	u.amountDol -= amountDol
}

func (u *User) addBit(amountDol float64) {
	u.amountDol += amountDol
}

func (u *User) removeBit(amountDol float64) {
	u.amountDol -= amountDol
}

func (u User) GetAmountDol() float64 {
	return u.amountDol
}

func (u User) GetAmountBit() float64 {
	return u.amountBit
}
