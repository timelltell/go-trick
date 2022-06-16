package main

import (
	"GolangTrick/Compare"
	"GolangTrick/Middle"
	"GolangTrick/tric"
	"fmt"
)

func TestCompare() {
	Compare.Compare()
}
func TestInterface() {
	d1 := &Compare.Country{"China"}
	d3 := Compare.Country{"USA"}
	d2 := Compare.City{"shenzhen"}
	Compare.PrintStr(d1)
	Compare.PrintStr(d2)
	fmt.Println(d3.ToString())
}

func TestComplete() {
	s := Compare.Square{4}
	fmt.Println(s.Sides())
}

func TestFunctional() {
	src, err := Compare.NewServer("127.0.0.1", 8080, Compare.Protocol("tcp"), Compare.Timeout(123))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("src: ", src)
}
func TestReduce() {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 4, 5}
	res := Compare.Map(nums, square)
	fmt.Println("res : ", res)

	type Employee struct {
		Name     string
		Age      int
		Vacation int
		Salary   int
	}
	var list = []Employee{
		{"Hao", 44, 0, 8000},
		{"Alice", 23, 5, 9000},
		{"Mike", 32, 8, 4000},
	}
	old := func(e Employee) bool {
		return e.Age > 40
	}
	res = Compare.Map(list, old)

	fmt.Printf("old people: %d\n", res)

}

func TestRedis() {
	Middle.RedisPrac()
}
func TestReflect() {
	tric.AllTypes()
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//time.Sleep(time.Second * 3)

	//fmt.Println("start")
	//tric.Testregexp()
	//TestCompare()
	//TestInterface()
	//TestComplete()
	//TestFunctional()
	//TestReduce()
	//TestRedis()
	//TestReflect()
	//tric.TestInject()
	//tric.TestInject1()
	//tric.TestInject2()

	//tric.TestSyncPool()
	//fmt.Println("middle")

	//mytime.TestTime()
	//fmt.Println("end")

	//my_select.Test1()
	//time.Sleep(time.Second * 100)

	//ch := make(chan *int, 2)
	//l := make([]int, 0, 2)
	//l = append(l, 1)
	//l = append(l, 2)
	////var a int = 3
	//for index, tmp := range l {
	//	fmt.Printf("%p\n", &tmp)
	//	fmt.Printf("%p\n", &l[index])
	//	test4(&l[index], ch)
	//}
	//fmt.Println("ch")
	//for i := range ch {
	//	fmt.Println(*i)
	//}

	functest()
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

func functest() {
	fmt.Println("a")
	fmt.Println(LimitStrategy0)
	fmt.Println("b")
	fmt.Println(LimitStrategy1)
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
