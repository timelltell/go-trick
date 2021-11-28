package tric

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type s1 interface {

}


type s2 interface {

}

func Formt(name s1, age s2){
	fmt.Printf("name=%s,age=%d",name,age)
}

type person struct {
	name string
	age int
}

func TestInject1()  {

	inj := inject.New()
	inj.MapTo("cb",(*s1)(nil))
	inj.MapTo(26,(*s2)(nil))
	inj.Invoke(func(name s1, age s2){
			fmt.Printf("name=%s,age=%d",name,age)
	})
}
func TestInject()  {
	p := &person{
		"cb-person",
		26,
	}

	inj := inject.New()
	inj.Map(p)
	inj.Invoke(func(p *person){
		fmt.Printf("name=%s,age=%d",p.name,p.age)
	})
}