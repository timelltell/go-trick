package my_select

import (
	"fmt"
	"time"
)

type person struct {
	Name string
}

func Test1() {
	ch := make(chan *person, 1)
	ch2 := make(chan *person, 1)
	tmp := &person{
		Name: "chen",
	}
	go func() {
		ch <- tmp
		ch2 <- tmp
	}()
	//time.Sleep(1 * time.Second)
	res := test(ch, ch2)
	fmt.Printf("person : %v\n", *res)
	go func() {
		for {
			res = <-ch
			fmt.Printf("res person : %v\n", *res)
		}
	}()
	go func() {
		for {
			res = <-ch2
			fmt.Printf("res2 person : %v\n", *res)
		}
	}()
	time.Sleep(1 * time.Second)
}

func test(ch chan *person, ch2 chan *person) *person {
	for {
		select {
		case res := <-ch:
			fmt.Println("res")
			return res
		case res := <-ch2:
			fmt.Println("res2")
			return res
			//default:
			//	fmt.Println("default")
			//	return &person{
			//		Name: "default",
			//	}
		}
	}
}
