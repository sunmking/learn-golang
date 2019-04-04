package main

import (
	"unsafe"
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
	intType()
	floatType()
	complexType()
	stringType()
	typeTurn()

	fmt.Println("类型")
}

//bool
func boolType(){
	a:=false
	b:=true

	//c:= b && a
	c:= b || a

	fmt.Println("a:",a,"b:",b,"c:",c)
}
//int
func intType(){
	var a uint32 = 89
	var c uint16 = 77
	b := 95

	fmt.Println("value of a is",a,"value of b is",a)
	fmt.Printf("\n type of a is %T, size of a is %d",a,unsafe.Sizeof(a))
	fmt.Printf("\n type of b is %T, size of b is %d",b,unsafe.Sizeof(b))
	fmt.Printf("\n type of c is %T, size of c is %d",c,unsafe.Sizeof(c))
}
//float
func floatType(){
	a,b:=2.22,3.55

	fmt.Printf("\n type of a is %T, size of c is %d",a,unsafe.Sizeof(a))
	fmt.Printf("\n type of b is %T, size of b is %d",b,unsafe.Sizeof(b))

	sum := a + b
	diff := b - a

	no1,no2 := 99,120
	fmt.Println("\n sum :",sum,"diff :",diff)
	fmt.Println(" sum :",no1+no2,"diff :",no1-no2)
}
//复数
func complexType(){
	c1 := complex(5,7)
	c2 := 8 + 27i

	cadd := c1 + c2

	fmt.Println("sum:",cadd)

	cmul := c1 * c2
	fmt.Println("product:",cmul)
}
//string
func stringType(){
	first:="alex"
	last:="sun"
	name := first + " " + last

	fmt.Println("name is",name)
}
//Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。我们通过一个例子说明这意味着什么。
func typeTurn(){
	i := 55      //int
    j := 67.8    //float64
    //sum := i + j //不允许 int + float64
    sum := i + int(j) //允许 int + int
	fmt.Println(sum)
	
	x := 10
    var jx float64 = float64(x) // 若没有显式转换，该语句会报错
    //var jx float64 = x // 没有显式转换，该语句报错
    fmt.Println("jx", jx)
}