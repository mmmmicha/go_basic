package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Println("a 는 10입니다.")
	}

	// if, else if
	a = 256
	if a == 200 {
		fmt.Println("a 는 200입니다.")
	} else if a == 256 {
		fmt.Println("a 는 256입니다")
	}
}
