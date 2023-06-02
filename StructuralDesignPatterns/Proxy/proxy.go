package proxy

import (
	"errors"
	"fmt"
)

type Bank struct {
	amount int
}

func (Bank *Bank) withDraw(amount int) error {
	if amount > Bank.amount {
		return errors.New("amount is too large")
	}
	Bank.amount -= amount
	fmt.Println("withdraw", amount)
	return nil
}
func (Bank *Bank) deposit(amount int) {
	Bank.amount += amount
	fmt.Println("deposit", amount)
}

func (Bank *Bank) check() {
	fmt.Println("amount=", Bank.amount)
}

type ProxyBank struct {
	Bank Bank
}

func (ProxyBank *ProxyBank) withDraw(amount int) error {
	error := ProxyBank.Bank.withDraw(amount)
	return error
}
func (ProxyBank *ProxyBank) deposit(amount int) {
	ProxyBank.Bank.deposit(amount)
}

func (ProxyBank *ProxyBank) check() {
	ProxyBank.Bank.check()
}

func Main() {
	bank := Bank{}
	proxyBank := ProxyBank{Bank: bank}

	proxyBank.deposit(1000)
	proxyBank.withDraw(50)
	proxyBank.check()
}
