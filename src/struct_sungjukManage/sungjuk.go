package main

import "fmt"

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) viewSungjuk() {
	fmt.Println(s.sungjuk)
}

// 아래의 두 inputSungjuk 함수가 같다고 보면됨
// 즉, 아무리 구조체를 생성해 두고 아래 첫번째 inputSungjuk을 실행해보면
// 객체 역시 함수에서 사용되는 일회용으로 쓰이기 때문에
// 이 값이 실제 객체에 까지 미치진 않는다.
// 객체를 return 해서 그 객체를 다시 받게된다면 모를까

func (s Student) inputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func inputSungjuk(s Student, name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func main() {

	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	s.viewSungjuk()

	// 진짜 중요한 부분
	s.inputSungjuk("과학", "A")
	s.viewSungjuk()
}
