package main

import "fmt"

func main() {
	//
	a := 0
	for a < 10 {
		fmt.Println("a=", a)
		a++
	}
	fmt.Println("최종값은", a)

	// Go에서의 스코프
	// i 를 지역변수로 선언해놔도 for에서 i 를 재선언하는 것이 가능하다.
	var i int
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("최종값은", i)

	// for {}
	// 무한루프
	// 기본적으로 for 와 {} 사이에 아무런 입력이 없으면 true 이다.

}
