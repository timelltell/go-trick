package myjson

import (
	"encoding/json"
	"fmt"
)
type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func Test(){
	str := "{\"name\":\"kimchenbin\",\"age\":16}"
	str2:= "{\"name\":\"kimchenbin\"}"
	var p person
	err:=json.Unmarshal([]byte(str),&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%+v",p)

	var p2 person
	err=json.Unmarshal([]byte(str2),&p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%+v",p2)
}
