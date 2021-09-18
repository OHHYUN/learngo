package main

import (
	"fmt"

	"github.com/OHHYUN/learngo/accounts"
)

// GO의 시작포인트
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
