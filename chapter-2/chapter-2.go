package main

import (
	"fmt"
)
//类型

/**
1、bool
2、数字类型
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
3、string
*/

func main(){
	boolType()

	fmt.Print("类型")
}

//bool
func boolType(){
	a:=false
	b:=true

	//c:= b && a
	c:= b || a

	fmt.Println("a:",a,"b:",b,"c:",c)
}