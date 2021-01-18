package main

import (
	"fmt"
)

// 变量

/**
声明单个变量
var name type 是声明单个变量的语法。
*/

func main() {
	//string
	//var str string = "我是变量"
	str := "我是变量"

	//int
	//var number int64 = 2
	number := 2

	//float
	//var float_data float64 = 3.0
	float_data := 3.0

	// 一行赋值
	str1, number1, float_data1 := "s1", 2, 4.0

	// var
	var (
		name   = "alex"
		sex    = 1
		height int
	)

	fmt.Print("变量：" + str)

	fmt.Printf("整型：%d", number)

	fmt.Printf("浮点：%.2f", float_data)

	fmt.Printf("s:%s d:%d f:%.2f", str1, number1, float_data1)

	fmt.Printf("s:%s d:%d f:%d", name, sex, height)
}
