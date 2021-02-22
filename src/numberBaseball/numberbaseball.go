package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 이게 무작위를 설정하는 것
	// 이걸 안할 경우 컴퓨터는 기본 default 로 1, 7, 9 ... 로 찍어낸다.
	rand.Seed(time.Now().UnixNano())
	// 1. 무작위 숫자 3개를 만든다
	numbers := MakeNumbers()
	fmt.Println(numbers)

	cnt := 0
	for {
		// 2. 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 3. 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 4. 3S 인가?
		if IsGameEnd(result) {
			// 5. 게임 끝
			break
		}
	}

	// 6. 게임 끝 몇번만에 끝났는지 출력
	fmt.Printf("%d번만에 맞췄습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicate := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicate = true
					break
				}
			}
			if !duplicate {
				rst[i] = n
				break
			}
		}
	}
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 비교 결과가 3 스트라이크 인지 확인
	return true
}
