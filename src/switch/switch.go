package main

import "fmt"

func main() {
	a := 10
	// Go에서의 switch는 break가 기본 내장되어있다.
	// java와 다르게 맨처음 조건이 맞으면 그 조건에서 결과를 얻고 break 한다.
	switch {
	case a <= 10:
		fmt.Println("a는 10보다 같거나 작습니다.")
	case a >= 10:
		fmt.Println("a는 10보다 같거나 큽니다")
	case a < 10:
		fmt.Println("a는 10보다 작습니다")
	}
}
