package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

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
		cnt++

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

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요")
		var no int
		// scanf를 쓸 때 \n를 안써주게되면 입력 후 엔터를 누르면 입력buffer에 \n가 남아있게 되어
		// 다음 엔터를 누를 때 엔터가 두번 눌러지는 현상이 발생한다.
		_, err := fmt.Scanf("%d\n", &no)

		// nil 무언가가 없는 경우
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		cnt := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicate := false
			for j := 0; j < cnt; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicate = true
					break
				}
			}
			if duplicate {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if cnt > 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[cnt] = n
			cnt++
		}

		if success && cnt < 3 {
			fmt.Println("3개의 숫자를 입력하세요")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	strikes := 0
	balls := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}
	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

func IsGameEnd(result Result) bool {
	// 비교 결과가 3 스트라이크 인지 확인
	return result.strikes == 3
}
