package tric

import "fmt"

//示例是否实现了某个方法，以及interface的赋值规则

type daughter struct {
	word string
}

//func (d *daughter) Say() {
//	fmt.Println("i am a daughter ", d.word)
//}
func (d *daughter) Say() {
	fmt.Println("i am a daughter ", d.word)
}

func test() {
	var a *daughter = new(daughter)
	_, ok := interface{}(a).(interface{ Say() })
	fmt.Println(ok)

	a.Say()

	var c daughter
	_, ok = interface{}(c).(interface{ Say() })
	fmt.Println(ok)

	c.Say()
}
