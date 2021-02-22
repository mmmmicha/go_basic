package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s Student) printSungjuk() {
	fmt.Println(s.class, s.grade)
}

// 매우중요!!
// 이렇게 structure 를 변수로 넘기는 경우엔
// 그 structure의 모든값들이 똑같이 복사되어 메모리에 적재되는데
// structure의 pointer를 변수로 넘기는 경우엔
// 그 structure의 메모리 주소만 복사되어 메모리에 적재되므로
// structure의 규모가 클 수록 pointer가 압도적으로 효과적이다.
func (s Student) inputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func (s *Student) printSungjuk_pointer() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) inputSungjuk_pointer(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var a int
	var p *int
	var b int

	p = &a
	a = 3
	b = 2

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)

	// 설명했던 변수의 복사 개념이기 때문에 변동 x
	increase(a)
	fmt.Println("그냥 함수 적용 시=", a)

	// 포인터로 넘기기
	increasePointer(&a)
	fmt.Println("pointer 함수 적용 시=", a)

	// 학생예시
	var s Student = Student{name: "광현", age: 28, class: "수학", grade: "A"}

	s.inputSungjuk("과학", "C")
	s.printSungjuk()

	s.inputSungjuk_pointer("과학", "C")
	s.printSungjuk_pointer()
}

func increase(x int) {
	x++
}

func increasePointer(x *int) {
	*x = *x + 1
}
