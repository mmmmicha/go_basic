package main

import "fmt"

// slice 는 동적배열을 의미한다.
// java의 ArrayList와 흡사하다고 보면 된다.

// 동적배열은 실제 배열을 포인트하고 있다.
// 필요한 사이즈가 늘어날때마다 2배씩 늘린다. ex. 시작이 3이었으면, 6으로 사이즈가 변함

// 아주 중요한 개념!! 메모리의 주소가 달라질 수 있기 때문에
// 1. 기존의 할당된 메모리 말고 2배가 늘어난 빈 배열을 다른 메모리에 생성한다.
// 2. 기존의 배열을 빈 배열에 복사한다.
// 3. 추가한 요소를 늘어난 배열쪽에 할당한다.

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

	b := []int{1, 2} // 길이는 0, capacity 2

	c := append(b, 3)

	// %p 는 주소 찍기
	fmt.Printf("%p %p\n", b, c)
	// b 와 c 의 주소가 다르다.
	// 앞서 얘기했던대로 사이즈가 늘어나면서 빈배열을 만들었기 때문
	// 만약 capacity 가 length를 담기에 충분했다면 똑같은 주소가 나올것

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d,", b[i])
	}
	fmt.Println()

	for i := 0; i < len(c); i++ {
		fmt.Printf("%d,", c[i])
	}
	fmt.Println()

	fmt.Println(cap(b), cap(c))
}
