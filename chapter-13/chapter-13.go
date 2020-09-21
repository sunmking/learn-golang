package main

import (
	"fmt"
)

func main() {
	//指针是一种存储变量内存地址（Memory Address）的变量。
	fmt.Println("指针")
	f1()
}

func f1() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		//指针的解引用可以获取指针所指向的变量的值。将 b 解引用的语法是 *b。
		//
		//通过下面的代码，可以看到如何使用解引用。
		fmt.Println("b after initialization is", b)
		fmt.Println("b after value is", *b)
		fmt.Println("b after value is", a)
		*b++
		fmt.Println("b after value is", *b)
	}
	// 向函数传递指针参数
	aval := 40
	fmt.Println(aval)
	change(&aval)
	fmt.Println(aval)

	// 不要向函数传递数组的指针，而应该使用切片

	str3 := [4]int{1, 2, 3, 4}

	modS(&str3)
	fmt.Println(str3)

	modS1(str3[:])
	fmt.Println(str3)

	//Go 不支持指针运算
	vcr := [...]int{1, 2, 3, 4}
	vce := &vcr
	//vce++
	//vce[0] = 10
	(*vce)[0] = 10
	fmt.Println(vce)

}
func modS(str *[4]int) {
	// (*str)[0]=6
	str[0] = 6
}

func modS1(str []int) {
	str[0] = 60
}
func change(val *int) {
	*val = 55
}
