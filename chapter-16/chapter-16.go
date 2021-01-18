package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, r := range ms {
		if r == 'a' || r == 'b' || r == 'c' {
			vowels = append(vowels, r)
		}
	}
	return vowels
}

// 接口的实际用途 start

type SalaryCalculator interface {
	SalaryCalc() int
}
type Contract struct {
	empId    int
	basicpay int
}
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) SalaryCalc() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) SalaryCalc() int {
	return c.basicpay
}
func toExp(s []SalaryCalculator) {
	var exp = 0

	for _, v := range s {
		exp = exp + v.SalaryCalc()
	}

	fmt.Printf("Total Expense Per Month $%d", exp)
}

// 接口的实际用途 end

// 接口的内部表示
type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

//空接口
func describeIf(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//类型断言
//类型断言用于提取接口的底层值（Underlying Value）。
//
//在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
//
//一段代码胜过千言。下面编写个关于类型断言的程序。
func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

//类型选择
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func findType2(i interface{}) {
	switch v := i.(type) {
	case SalaryCalculator:
		fmt.Printf("I am %d\n", v.SalaryCalc())
	default:
		fmt.Printf("Unknown type\n")
	}
}

// 指针&值 引用多接口实现

type AgeInterface interface {
	showAge()
}
type SalaryInterface interface {
	showSalary()
}

type People struct {
	age    int
	salary int
}

func (p *People) showAge() {
	fmt.Println(p.age)
}
func (p People) showSalary() {
	fmt.Println(p.salary)
}

// 接口嵌套
type PeopleOp interface {
	AgeInterface
	SalaryInterface
}

// 接口的0值
type Describer interface {
	Describe()
}

func main() {
	fmt.Println("接口")
	ms := MyString("abc")
	var v VowelsFinder
	v = ms
	fmt.Printf("%c", v.FindVowels())

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	toExp(employees)
	//
	mf := MyFloat(3.44)
	var t Test
	t = mf
	describe(t)
	t.Tester()
	//
	s := string("sss")
	describeIf(s)

	i := 55
	describeIf(i)

	x := struct {
		name int
	}{
		name: 44,
	}
	describeIf(x)
	//类型断言
	var ss1 interface{} = 56.0
	assert(ss1)
	var ss interface{} = 99
	assert(ss)

	findType(ss1)
	findType(ss)

	//接口
	xc := "nab"
	findType2(xc)
	pm := Permanent{
		1,
		3,
		4,
	}
	findType2(pm)
	// 多接口实现
	p := People{
		20,
		3000,
	}
	var ag AgeInterface
	ag = &p
	ag.showAge()
	var xd SalaryInterface
	xd = &p
	xd.showSalary()
	//接口嵌套
	var pop PeopleOp = &p
	pop.showAge()
	pop.showSalary()

	// 接口的0值
	var de Describer
	if de == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", de, de)
	}
}
