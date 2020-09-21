package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。让我们来看一个创建并打印字符串的简单示例。
	fmt.Println("字符串")
	f1()
}

func f1() {
	st := "hello world"

	for i := len(st) - 1; i >= 0; i-- {
		fmt.Printf("value is %c \n", st[i])
	}

	for i := 0; i < len(st); i++ {
		fmt.Printf("value is %c \n", st[i])
	}
	//rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，
	//rune 表示一个代码点。代码点无论占用多少个字节，
	//都可以用一个 rune 来表示。让我们修改一下上面的程序，
	//用 rune 来打印字符。

	c := "ceshi"

	tc := []rune(c)

	for i := 0; i < len(tc); i++ {
		fmt.Printf("value is %c \n", tc[i])
	}

	for index, runee := range tc {
		fmt.Printf("index is %d value is %c \n", index, runee)
	}
	//用字节切片构造字符串
	vst := []byte{67, 97, 102, 195, 169}

	str := string(vst)

	fmt.Println(str)

	//用 rune 切片构造字符串

	vst1 := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str1 := string(vst1)
	fmt.Println(str1)

	// 字符串的长度
	length(str1)

	//字符串是不可变的

	xs := changstr([]rune("uususu"))

	fmt.Println(xs)

}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func changstr(s []rune) string {
	s[0] = 'a'
	return string(s)
}
