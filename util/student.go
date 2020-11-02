package util

type Student struct {
	Name string
	Age  int
}

func (t *Student) Newname(n string) {
	t.Name = n
	// fmt.Println(t)
}

// func (t Student) Newname(n string) {
// 	t.Name = n
// }
