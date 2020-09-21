package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

type StorageStock struct {
	id          int
	batch_no    string
	order_no    string
	quantity    float64
	create_time int
}
type Person struct {
	string
	int
	test int
}

type Person1 struct {
	name    string
	address Address
}

type Person2 struct {
	name string
	Address
}
type Address struct {
	city, state string
}

func main() {
	//结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
	fmt.Println("结构体")
	f1()
}

func f1() {
	emp := Employee{
		firstName: "join",
		lastName:  "sun",
		age:       20,
	}

	emp1 := Employee{"ha", "yun", 20, 100}

	fmt.Println(emp)
	fmt.Println(emp1)

	//匿名结构体

	x9587 := struct {
		name string
		age  int
	}{
		"xxx",
		20,
	}
	fmt.Println(x9587)
	fmt.Println(x9587.name)

	ss := StorageStock{1, "RK00000001", "ZRK0000003", 30, 160000000}

	fmt.Println(ss)
	//结构体 零值
	var dd StorageStock
	fmt.Println(dd)

	// 指针

	emp3 := &Employee{
		firstName: "HHH",
		lastName:  "DDD",
		age:       0,
		salary:    0,
	}
	fmt.Println((*emp3).lastName)
	fmt.Println(emp3.lastName)
	fmt.Println(emp3)
	fmt.Println(&emp3)

	//匿名字段

	ps := Person{"test", 3, 4}
	fmt.Println(ps.string)
	fmt.Println(ps.int)
	fmt.Println(ps.test)

	// 嵌套结构体

	ps1 := Person1{name: "xxx"}
	ps1.address = Address{"shanghai", "CN"}

	fmt.Println(ps1)
	//fmt.Println(ps1.city)

	// 提升字段
	ps3 := Person2{
		name: "kun",
	}

	ps3.Address = Address{
		city:  "shang",
		state: "ZH",
	}
	fmt.Println(ps3)
	fmt.Println(ps3.city)
}
