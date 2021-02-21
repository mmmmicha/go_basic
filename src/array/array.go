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
	// go에서의 string
	// String 은 []byte 기본배열로 나타낼수도 있지만, []rune 배열로 나타낼 수 있다.
	// []rune : 데이터에 다라 1~3byte 할당
	// s2 := []rune(s)
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
