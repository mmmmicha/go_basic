package main

import "fmt"

// slice를 slice 해보자

// 매우 중요한 부분은, slice 의 slice 는 같은 메모리 영역이다!!

// a[4:8] 은 4번째부터 7번째까지를 자른다는 의미
// a[4: ] 은 4번째부터 끝까지 자른다
// a[:4] 은 처음부터 3번째까지 자른다
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	c := a[4:]
	d := a[:4]

	fmt.Println(b)

	b[0] = 1
	b[1] = 2

	fmt.Println(b)
	// b를 바꿨는데도 c와 d가 영향을 받은것을 볼 수 있음.
	fmt.Println(c)
	fmt.Println(d)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		var back int
		// a, back = RemoveBack(a)
		a, back = RemoveFront(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println()
	fmt.Println(a)
}

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
