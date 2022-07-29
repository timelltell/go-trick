package my_trick

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type s1 interface {
}

type s2 interface {
}

func Formt(name s1, age s2) {
	fmt.Printf("name=%s,age=%d", name, age)
}

type person struct {
	name string
	age  int
}

func TestInject1() {

	inj := inject.New()
	inj.MapTo("cb", (*s1)(nil))
	inj.MapTo(26, (*s2)(nil))
	inj.Invoke(func(name s1, age s2) {
		fmt.Printf("name=%s,age=%d\n", name, age)
	})
}
func TestInject() {
	p := &person{
		"cb-person",
		26,
	}

	inj := inject.New()
	inj.Map(p)
	inj.Invoke(func(p *person) {
		fmt.Printf("name=%s,age=%d\n", p.name, p.age)
	})
}
func TestInject2() {
	p := &person{
		"cb-person",
		26,
	}

	inj := inject.New()
	inj.Map(p)
	inj.Map(1)
	inj.Map(2)
	inj.Map("s")
	inj.Map("s2")
	i, err := inj.Invoke(func(p *person, i int, s string) (string, int) {
		fmt.Printf("name=%s,age=%d, i=%d, s=%s\n", p.name, p.age, i, s)
		return s, i
	})
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Printf("i, %+v\n", i)
	fmt.Printf("i, %v\n", i[0])
	fmt.Printf("i, %d\n", i[1])
	fmt.Printf("i %v", 2)
}
