package main

import (
	"GolangTrick/Middle"
	"GolangTrick/my_trick"
	"fmt"
)

func TestRedis() {
	Middle.RedisPrac()
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//i := functest()
	//fmt.Println("i")
	//fmt.Println(i)
	//err := errors.New("")
	//if err != nil {
	//	fmt.Println("yes")
	//
	//}
	str := my_trick.GetFormatStr()
	fmt.Println(str)
}

func test4(i *int, ch chan<- *int) (res string) {
	fmt.Printf("%p\n", i)

	fmt.Println(*i)
	ch <- i
	return res
}

//func test2(m interface{}) {
//	m1, ok := m.(map[string][]int)
//	fmt.Println(ok)
//	for k, v := range m1 {
//		fmt.Println(k)
//		fmt.Println(v)
//	}
//}

func test3() (res string) {
	defer func() {
		res = "2"
	}()
	res = "1"
	return res
}

func test(list *[]person) {
	(*list)[0].Name = "kim"
	(*list) = append([]person{person{
		"kimchenbin2",
		27,
	}}, *list...)
}

const (
	a = iota
	b
)

const (
	LimitStrategy0 = LimitStrategyType(0) //表示每个请求参数，只请求一次下游的限流策略
	LimitStrategy1 = LimitStrategyType(1) //表示每个请求参数，只请求一次下游的限流策略
)

type ResProcessLimitDto struct {
	Request map[string]interface{}
}

type LimitStrategyType int

func functest() (status int) {
	defer func() {
		fmt.Println(status)
	}()
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for tmp := range ch {
		fmt.Println(tmp)
	}
	//fmt.Println("a")
	//fmt.Println(LimitStrategy0)
	//fmt.Println("b")
	//fmt.Println(LimitStrategy1)
	return -1
}

func funa() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("a %v\n", err)
		}
	}()
	funb()
	fmt.Println("funa")
}

func funb() {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	}
	//}()
	fmt.Println("funb")
	panic("panic")
}

//https://junedayday.github.io/categories/
