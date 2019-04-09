package main

import (
	"fmt"
)

/**
switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 if else 子句的常用方式。
*/
func main() {
	fmt.Println("switch语句")
	f1()
	f2()
	f3()
	f4()
}

// 默认
func f1() {
	// fi := 3
	// switch fi {
	// case 1:
	// 	fmt.Println("test ", 1)
	// case 2:
	// 	fmt.Println("test ", 2)
	// case 3:
	// 	fmt.Println("test ", fi)
	// default:
	// 	fmt.Println("no match")
	// }

	switch fi := 3; fi {
	case 1:
		fmt.Println("test ", 1)
	case 2:
		fmt.Println("test ", 2)
	case 3:
		fmt.Println("test ", fi)
	default:
		fmt.Println("no match")
	}
}

// 多表达式判断
func f2() {
	fi := "i"
	switch fi {
	case "a", "b":
		fmt.Println("match")
	default:
		fmt.Println("no match")
	}
}

// 无表达式的 switch
func f3() {
	num := 4
	switch {
	case num > 2:
		fmt.Println(true)
	default:
		fmt.Println(false)
	}
}
func number() int {
	num := 15 * 5
	return num
}

//Fallthrough 语句
func f4() {

	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

}
