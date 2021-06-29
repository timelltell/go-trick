package main

import (
	"GolangTrick/Compare"
	"fmt"
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
func main(){
	//TestCompare()
	//TestInterface()
	//TestComplete()
	TestFunctional()
}

