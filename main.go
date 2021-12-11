package main

import (
	"GolangTrick/Compare"
	"GolangTrick/Middle"
	"GolangTrick/tric"
	"fmt"
	"time"
)

func TestCompare(){
	Compare.Compare()
}
func TestInterface(){
	d1:=&Compare.Country{"China"}
	d3:=Compare.Country{"USA"}
	d2:=Compare.City{"shenzhen"}
	Compare.PrintStr(d1)
	Compare.PrintStr(d2)
	fmt.Println(d3.ToString())
}

func TestComplete(){
	s:=Compare.Square{4}
	fmt.Println(s.Sides())
}

func TestFunctional(){
	src,err :=Compare.NewServer("127.0.0.1",8080,Compare.Protocol("tcp"),Compare.Timeout(123))
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("src: ",src)
}
func TestReduce(){
	square := func(x int) int{
		return x*x
	}
	nums:=[]int{1,2,4,5}
	res:=Compare.Map(nums,square)
	fmt.Println("res : ",res)

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
	res=Compare.Map(list,old)

	fmt.Printf("old people: %d\n", res)

}

func TestRedis(){
	Middle.RedisPrac()
}
func TestReflect(){
	tric.AllTypes()
}
type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	//TestCompare()
	//TestInterface()
	//TestComplete()
	//TestFunctional()
	//TestReduce()
	//TestRedis()
	//TestReflect()
	//tric.TestInject()
	//tric.TestSyncPool()
	//mytime.TestTime()
	var list = make([]person,0,20)
	list = append(list, person{
		"kimchenbin",
		26,
	})
	test(&list)
	fmt.Println("list",list)

	startTime := time.Now()
	fmt.Println(startTime)
	time.Sleep(time.Second)
	fmt.Println(time.Since(startTime).Nanoseconds()/time.Millisecond.Nanoseconds())



}
func test(list *[]person)  {
	(*list)[0].Name="kim"
	(*list) = append([]person{person{
		"kimchenbin2",
		27,
	}},*list...)
}
//https://junedayday.github.io/categories/

