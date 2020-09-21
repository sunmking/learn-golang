package main

import (
	"../rectangle"
	"fmt"
	"log"
)

/**
包的初始化顺序如下：

首先初始化包级别（Package Level）的变量
紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。
如果一个包导入了另一个包，会先初始化被导入的包。

尽管一个包可能会被导入多次，但是它只会被初始化一次。
*/

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("包")
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
