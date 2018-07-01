package model

import "github.com/agiledragon/ddd-sample-in-golang/cargo/domain/model/base"

type balance struct {
	base.ValueObject
	amount uint
}

func newBalance(amount uint) *balance {
	return &balance{amount: amount}
}

func (this *balance) increase(amount uint) {
	this.amount += amount
}

func (this *balance) decrease(amount uint) {
	this.amount -= amount
}

func (this *balance) get() uint {
	return this.amount
}
