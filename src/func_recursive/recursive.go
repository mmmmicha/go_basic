package main

import "fmt"

func main() {

	fun1(10)

}

// 쭉 깊이 들어가서 밑에서부터 결과를 하나씩 리턴하면서 돌아오는 recursive 구조
// 모든 재귀호출은 반복문으로 바꿀 수 있다.
// 재귀호출이 반복문보다 느리다
// 재귀호출이 더 유리한 경우 : 피보나치 수열

func fun1(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	fun1(x - 1)
}
