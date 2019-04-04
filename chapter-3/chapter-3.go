package main

import (
	"math"
	"fmt"
)

func main()  {
	fmt.Println("常量")
	f1()
	f2()
	f3()
}
//常量不能再重新赋值为其他的值
func f1(){
	const a = 55 // 允许
    //a = 89       // 不允许重新赋值
}
//常量的值会在编译的时候确定。
//因为函数调用发生在运行时，
//所以不能将函数的返回值赋值给常量。
func f2(){
	fmt.Println("hello, playground")
	var a = math.Sqrt(4) //允许
	//const b = math.Sqrt(4) //不允许

	fmt.Println(a)
	//fmt.Println(b)
}
//Go 是一个强类型的语言，在分配过程中混合类型是不允许的。
//让我们通过以下程序看看这句话是什么意思。
func f3(){
	var defaultName = "Sam1" // 允许
	type myString string
	var customName myString = "Sam2" // 允许
	//customName = defaultName // 不允许
	customName = myString(defaultName) // 允许

	fmt.Println("dn:",defaultName,"cn:",customName)
}