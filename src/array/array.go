package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	s = "Hello 월드"
	// 인코딩 한글 3byte

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ", ")
	}
	// H, e, l, l, o,  , ì, , , ë, , ,
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}
	// 72, 101, 108, 108, 111, 32, 236, 155, 148, 235, 147, 156,
	fmt.Println("len(s)의 길이=", len(s))
	// len(s)의 길이= 12

	// go에서의 string
	// String 은 []byte 기본배열로 나타낼수도 있지만, []rune 배열로 나타낼 수 있다.
	// []rune : 데이터에 다라 1~3byte 할당
	// s2 := []rune(s)
	s2 := []rune(s)
	fmt.Println("len(s2)=", len(s2))
	// len(s2)= 8
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
	// H, e, l, l, o,  , 월, 드,
	fmt.Println()
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
	// 72, 101, 108, 108, 111, 32, 50900, 46300,
}
