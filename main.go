package main

import (
	"fmt"
	"strconv"

	"test/example/util"
)

// 인터페이스 구현할 스트럭처
type StructA struct {
}

// 포인터 리시버
//
//	func (sa *StructA) AAA(a int) int {
//		return a
//	}
//
//	func (sa *StructA) BBB(a int) string {
//		return "a=" + strconv.Itoa(a)
//	}
//
// 값 리시버
func (sa StructA) AAA(a int) int {
	return a
}
func (sa StructA) BBB(a int) string {
	return "a=" + strconv.Itoa(a)
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("안녕, 내 이름은 ", p.Name, "야")
}

type Android struct {
	Person
	Model string
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	st := util.Student{Name: "st-name", Age: 10}
	//소는 누가 키우나 누가할거냐지 추제와 동사 목적어 c.Newname("haha") c라는 객체가 이름을 바꾼다 haah로
	//go 에서는 파라미터는 무조건 카피된다 그러나 주소가 카피되는냐 값이 카피되느냐의 차이일뿐
	// var c *util.Student = &util.Student{Name: "hahaha", Age: 20}
	var c util.Student = util.Student{Name: "c-name", Age: 20}
	d := c
	fmt.Println(c)
	c.Newname("c-newname")
	d.Newname("d-newname")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(st)
	/////////////////////////////////////////////
	//인터페이스 테스트 포인터와 값
	// var i util.InterfaceA = &StructA{}
	var i util.InterfaceA = StructA{}
	fmt.Println(i.AAA(2))
	fmt.Println(i.BBB(3))
	////////////////////////////////////////////
	xPtr := new(int) //new 함수를 이용해 포인터 변수 생성
	one(xPtr)
	fmt.Println(*xPtr) //  1
	///////////////////////////////////////////
	// is-a 관계는 이처럼 직관적으로 동작한다. 즉, 사람은 이야기할 수 있고,
	// 안드로이드는 사람이며, 따라서 안드로이드는 이야기할 수 있다.
	// a := new(Android)
	// a.Person.Talk()
	a := new(Android)
	a.Talk()
}
