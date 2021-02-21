package main

import "fmt"

func main() {

	// 스타성 찍기
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 5-2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
