package main

import (
	"fmt"

	"github.com/OHHYUN/learngo/accounts"
)

// GO의 시작포인트
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	fmt.Println(account)
}
