package main

import (
	"fmt"
	"golang/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	account.Deposit(10)
	
	fmt.Println(account.Balance())
}