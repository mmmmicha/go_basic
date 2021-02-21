package main

import "fmt"

// 역배열 만들어보기
// 방식 2개
func main() {
	// 방식 1
	arr := [5]int{1, 2, 3, 4, 5}
	// arr := [5]{1, 2, 3, 4, 5}
	// 이런형태는 x
	clone := [5]int{}
	temp := [len(arr)]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
		temp[i] = arr[len(arr)-1-i]
	}
	fmt.Println(clone)
	fmt.Println(temp)

	// 방식 2
	// go에서 가능한 이중대입
	// 일단 앞으로 보일 예제는 java에서 xor 를 대체할 수 있음
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr)
}
