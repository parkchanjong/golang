package main

import (
	"fmt"
	"golang/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	fmt.Println(account);
}