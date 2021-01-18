package main

import (
	"fmt"
	"math"
)

type People struct {
	name string
	age  int
}

//方法其实就是一个函数，
//在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
//接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的

//方法
func (p People) display() {
	fmt.Println(p.name)
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e People) {
	fmt.Printf("Salary of %s is %d", e.name, e.age)
}

//
type RectAngle struct {
	width  float64
	length float64
}
type Circle struct {
	radius float64
}

func (r RectAngle) Area() float64 {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//指针接收器与值接收器
//到目前为止，我们只看到了使用值接收器的方法。
//还可以创建使用指针接收器的方法。
//值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。让我们用下面的程序来帮助理解这一点。

func (p People) changeName(newName string) {
	p.name = newName
}

func (p *People) changeAge(newAge int) {
	p.age = newAge
}

//当一个函数有一个值参数，它只能接受一个值参数。
//当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
//让我们通过一个例子来理解这一点。

func area(r RectAngle) float64 {
	return r.width * r.length
}

func (r *RectAngle) area() float64 {
	return r.length * r.width
}

//我们尝试把一个 add 方法添加到内置的类型 int。
//这是不允许的，因为 add 方法的定义和 int 类型的定义不在同一个包中。
//该程序会抛出编译错误 cannot define new methods on non-local type int。
//让该程序工作的方法是为内置类型 int 创建一个类型别名，然后创建一个以该类型别名为接收器的方法。
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}
func main() {
	fmt.Printf("函数")

	pp := People{
		name: "333",
		age:  777,
	}
	pp.display()
	displaySalary(pp)

	rec := RectAngle{
		length: 2.00,
		width:  4.00,
	}
	fmt.Printf("\r\n %.2f", rec.Area())

	cir := Circle{
		radius: 20.4,
	}
	fmt.Printf("\r\n %.2f", cir.Area())

	pp.changeName("6666")
	(&pp).changeAge(33)

	fmt.Printf("\r\n %s  %d", pp.name, pp.age)

	//z :=rectangle.Area(20.2,20.4)
	//
	//fmt.Println(z)

	rx := RectAngle{
		length: 3.4,
		width:  4.4,
	}

	fmt.Printf("\r\n %.2f", (&rx).area())
	fmt.Printf("\r\n %.2f \r\n", area(rx))

	a := myInt(2)
	b := myInt(4)

	c := a.add(b)

	fmt.Println(c)

}
