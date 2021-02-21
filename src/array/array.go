package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	s = "Hello 월드"
	// H, e, l, l, o,  , ì, , , ë, , ,
	// 인코딩 한글 3byte

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}
	fmt.Println("len(s)의 길이=", len(s))

	s2 := []rune(s)
	fmt.Println("len(s2)=", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
	fmt.Println()
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
}
