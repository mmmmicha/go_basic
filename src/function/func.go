package main

import (
	"fmt"
)

// go 는 참조호출, 복사호출 중에 무조건 값이 복사되는 특성을 가지고 있다.
// 즉, 새로운 주소에 값을 갖게 되는 것이다.

// 응집도 높이고, 종속성 낮추고
func add(x int, y int) int {
	return x + y
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a, b := fun1(2, 3)
	fmt.Println(a, b)
}

func fun1(x, y int) (int, int) {
	//return y, x
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println("func2", x, y)
}
