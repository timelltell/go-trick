package tric

import (
	"fmt"
	"reflect"
)

func AllTypes(){
	for i:=reflect.Invalid;i<=reflect.UnsafePointer;i++{
		fmt.Println("i: ",i.String())
	}
}
