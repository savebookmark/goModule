package main

import (
	"fmt"
	"strconv"

	"github.com/savebookmark/testebiten/util"
)

//인터페이스 구현할 스트럭처
type StructA struct {
}

//포인터 리시버
// func (sa *StructA) AAA(a int) int {
// 	return a
// }
// func (sa *StructA) BBB(a int) string {
// 	return "a=" + strconv.Itoa(a)
// }
//값 리시버
func (sa StructA) AAA(a int) int {
	return a
}
func (sa StructA) BBB(a int) string {
	return "a=" + strconv.Itoa(a)
}

func main() {
	st := util.Student{Name: "HIHI", Age: 10}
	//소는 누가 키우나 누가할거냐지 추제와 동사 목적어 c.Newname("haha") c라는 객체가 이름을 바꾼다 haah로
	//go 에서는 파라미터는 무조건 카피된다 그러나 주소가 카피되는냐 값이 카피되느냐의 차이일뿐
	// var c *util.Student = &util.Student{Name: "hahaha", Age: 20}
	var c util.Student = util.Student{Name: "hahaha", Age: 20}

	// d := c

	fmt.Println(c)

	c.Newname("hohohoh")
	// d.Newname("dha")

	fmt.Println(c)
	// fmt.Println(d)

	fmt.Println(st)
	/////////////////////////////////////////////
	//인터페이스 테스트 포인터와 값
	// var i util.InterfaceA = &StructA{}
	var i util.InterfaceA = StructA{}

	fmt.Println(i.AAA(2))
	fmt.Println(i.BBB(3))
}
