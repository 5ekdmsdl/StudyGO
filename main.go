package main

import (
	accounts "GoLang/banking"
	"fmt"
	"log"
)

func main() {
	account := accounts.NewAccount("daeuna")
	account.Deposit(10)
	println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	println(account.Balance())
}
