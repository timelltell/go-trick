package Compare

import (
	"fmt"
	"reflect"
)

func Compare(){
	a:=1
	b:=1
	slice1:=[]*int{nil,&a}
	slice2:=[]*int{nil,&b}
	m1:=make(map[string]interface{})
	m1["a"]="a"
	m1["b"]=&b
	m2:=make(map[string]interface{})
	m2["a"]="a"
	m2["b"]=&a
	fmt.Println("v1 == v2: ", reflect.DeepEqual(slice1,slice2))
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1,m2))
}