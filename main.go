package main

import accounts "GoLang/banking"

func main() {
	account := accounts.NewAccount("daeuna")
	account.Deposit(10)
	println(account.Balance())
}
