package main

import (
	"fmt"
)

/**
for initialisation; condition; post {
}
*/
func main() {
	fmt.Println("循环")
	f1()
	f2()
	f3()
	f4()
	//f5()
}

//break
//break 语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。
func f1() {
	for num := 0; num <= 20; num++ {
		fmt.Println(num)
		if num > 10 {
			break
		}
	}
	fmt.Println("end 1")
}

//continue
//continue 语句用来跳出 for 循环中当前循环。
//在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。
//循环体会在一下次循环中继续执行。
func f2() {
	for num := 0; num <= 100; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Println(num)
	}
	fmt.Println("end 2")
}

//
func f3() {
	for n := 0; n <= 10; {
		fmt.Printf("\n%d", n)
		n += 1
	}
}

func f4() {
	n := 0
	for n <= 10 {
		fmt.Printf("\n%d", n)
		n += 1
	}
}

func f5() {
	for {
		fmt.Println("无限循环")
	}
}
