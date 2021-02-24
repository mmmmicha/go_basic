package main

import "fmt"

// slice 는 동적배열을 의미한다.
// java의 ArrayList와 흡사하다고 보면 된다.

// 동적배열은 실제 배열을 포인트하고 있다.
// 필요한 사이즈가 늘어날때마다 2배씩 늘린다.

// capacity : 할당되어있는 사이즈
// length : 현재 사용되고 있는 길이
// * 2의배수로 사이즈가 느는경우는 동적으로 사이즈가 오바될때임을 참고!@
// ex) 3개를 사용하고 싶으면 2의배수로 사이즈가 늘기때문에 capacity:4, length:3
// ex) 5개를 사용하고 싶으면 2의배수로 사이즈가 늘기때문에 capacity:8, length:5

func main() {
	// var a []int
	// a := []int{1, 2, 3, 4, 5}
	a := make([]int, 0, 8) // 길이는 0, capacity 8
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	a = append(a, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
}
