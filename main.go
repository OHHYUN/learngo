package main

import "fmt"

// GO의 시작포인트
func main() {
	urls := []string{
		"https://www.airbnb.com/",
	}

	for _, url := range urls {
		fmt.Println(url)
	}
}

func hitURL(url string) {

}
