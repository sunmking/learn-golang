package main

import (
	"fmt"
)

func main()  {
	fmt.Println("函数")

	fmt.Println(dealNum(2,3))
	fmt.Println(dealNum1(3,5))
	
	//_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
	x,_:=dealNum1(3,7)

	fmt.Println(x)
}

func dealNum(a,b int)(int,int){
	sum := a + b
	mul := a * b
	return sum,mul
}
//从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
func dealNum1(a,b int)(sum,mul int){
	sum = a + b
	mul = a * b
	return // 不需要明确指定返回值，默认返回sum,mul 的值
}