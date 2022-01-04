package main

import (
	"GolangTrick/Compare"
	"GolangTrick/Middle"
	"GolangTrick/tric"
	"GolangTrick/tric/mytime"
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
	//var list = make([]person, 0, 20)
	//list = append(list, person{
	//	"kimchenbin",
	//	26,
	//})
	//fmt.Println("list", &list)
	//fmt.Printf("%p\n", list)
	//
	//test(&list)
	//fmt.Println("list", &list)
	//fmt.Printf("%p\n", list)
	//
	//startTime := time.Now()
	//fmt.Println(startTime)
	//time.Sleep(time.Second)
	//fmt.Println(time.Since(startTime).Nanoseconds() / time.Millisecond.Nanoseconds())
	//f := test
	//funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	//fmt.Println(funcName)
	//ifData1 := 1
	//var ifData interface{} = ifData1
	//
	//switch realData := ifData.(type) {
	//case json.Number:
	//	_, _ = realData.Float64()
	//	fmt.Println("json number")
	//case float64:
	//	fmt.Println("json number4")
	//
	////result = realData
	//case float32:
	//	fmt.Println("json number1")
	//
	////result = float64(realData)
	//case int:
	//	fmt.Println("json number2")
	//
	////result = float64(realData)
	//case int64:
	//	fmt.Println("json number3")
	//
	////result = float64(realData)
	//default:
	//	fmt.Println("real ")
	//	fmt.Println(realData)
	//	//err = errors.New(fmt.Sprintf("%T to float64 is not support", ifData))
	//}

	//var p person = person{
	//	"kimchenbin",
	//	26,
	//}
	//byteses, _ := json.Marshal(&p)
	//str1 := string(byteses)
	//_ = json.Unmarshal([]byte(str1), &p)
	//byteses2, _ := json.Marshal(&p)
	//str2 := string(byteses2)
	//fmt.Println(str2)
	//fmt.Println(str1)
	//fmt.Println(str2 == str1)
	//var p2 person = person{
	//	"kimchenbin",
	//	27,
	//}
	//type fa struct {
	//	Id int
	//	P  *person
	//}
	//var f1 fa = fa{
	//	1,
	//	&p,
	//}
	//var f2 fa = fa{
	//	1,
	//	&p2,
	//}
	//fmt.Println(reflect.DeepEqual(p, p2))
	//fmt.Println(reflect.DeepEqual(&f1, &f2))

	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer func() {
	//			wg.Done()
	//			err := recover()
	//			if err != nil {
	//				fmt.Println("err")
	//
	//			}
	//		}()
	//		if i > 8 {
	//			return
	//		}
	//
	//		fmt.Println("i: ", i)
	//		//wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//fmt.Println("success: ")

	fmt.Println(mytime.CompareTwoDateTime("2022-01-04 20:11:49", "2022-01-03 20:11:49"))

}
func test(list *[]person) {
	(*list)[0].Name = "kim"
	(*list) = append([]person{person{
		"kimchenbin2",
		27,
	}}, *list...)
}

//https://junedayday.github.io/categories/
