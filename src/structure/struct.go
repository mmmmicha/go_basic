// 구조체(Structure) : 객체와 비슷
// 개념을 한곳에 모아둔 것
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "말똥이"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)
	// { 0} {개똥이 15} {소똥이 21} {말똥이 0} { 0}
	// Person 을 단순 선언했을 때 초기값이 0 이다.
	// null 이나 undefined 개념이 어떻게 되어있는지 확인해볼 필요 있음.

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.PrintName()
}
