package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t Student, newName string) {
	t.name = newName
}

func SetName_pointer(t *Student, newName string) {
	t.name = newName
}

// 이렇게 func 앞에 structure 가 명시되어 있는 경우는
// method 라 칭한다.
// func 와 method 의 차이!!
func (t *Student) SetName_method(newName string) {
	t.name = newName
}

func (t *Student) SetAge_method(newAge int) {
	t.age = newAge
}

func (t *Student) SetGrade_methode(newGrade int) {
	t.grade = newGrade
}

func main() {
	a := Student{"aaa", 20, 10}

	// := 는 copy
	// b := a

	// reference 타입의 관계
	var b *Student
	b = &a
	b.age = 30

	fmt.Println(a)
	fmt.Println(b)
	// value 타입에 대한 대입
	// 복사만 하고 그 복사한 놈의 데이터를 바꾼것
	SetName(a, "bbb")
	fmt.Println(a)

	// reference 타입으로 파라미터 넘김
	// 원본에 변화가 생김
	SetName_pointer(b, "bbb")
	fmt.Println(a)

	// oop 의 개념 -> 절차중심에서 object중심으로의 변화
	// 주어 동사 목적어 --> 주어.동사(목적어) / 주어 : object, 목적어 : object / 이 둘의 relationship
	// Entity의 relationship 이 중요해졌다!! ER
	// instance 개념 -> 생명주기를 부여하는 것
	// method 의 사용 -> 패러다임의 변화
	a.SetName_method("ccc")
	fmt.Println(a)

	var c *Student
	c = &Student{"ccc", 20, 10}

	c.SetName_method("ddd")
	c.SetAge_method(40)
	c.SetGrade_methode(40)
	fmt.Println(c)
}
