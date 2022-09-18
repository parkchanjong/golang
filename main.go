package main

import (
	"fmt"
	"golang/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	account.Deposit(10)

	fmt.Println(account)
	
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
} 