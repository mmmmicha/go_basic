// radix sort
// 주의사항
// 1. 0 ~ n n의 값이 정해져있어야함
// 2. 범위의 갯수가 적어야 함

package main

import "fmt"

func main() {
	arr := [10]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4}
	// 참고
	// go 의 변수할당 시 default 값은 항상 있음. int 같은 경우 0

	// 이처럼 사이즈를 10으로 잡을 수 있는 이유는 우리가 최대값을 9로 잡고
	// 테스트를 하고 있기 때문임
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}
