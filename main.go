package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

// GO의 시작포인트
func main() {
	superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	//축약형은 함수 내에서만 가능

}
