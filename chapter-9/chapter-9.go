package main

import "fmt"

/**
数组和切片
数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。
（译者注：当然，如果是 interface{} 类型数组，可以包含任意类型）
*/
func main() {
	fmt.Println("数组和切片")
	//数组
	f1()
	f2()
	f3()
	f4()
	f5()
	//切片
	f6()
	//
	f7()
	f8()
}

func f1() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	//fmt.Println(a)
	fmt.Println(a[:1])

	var b = [3]int{2, 3, 4}

	fmt.Println(b)

	c := [3]int{3, 4, 5}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}

	fmt.Println(d)
}

//数组是值类型
func f2() {
	a := [2]int{1, 3}
	var b [5]int

	//b = a //cannot use a (type [2]int) as type [5]int in assignmentgo
	fmt.Println(a, b)
}

//for 循环可用于遍历数组中的元素。
func f3() {
	var a = [...]float64{1.2, 2.1, 3.4}

	fmt.Println("len", len(a), "cap", cap(a))

	var le = len(a)

	for i := 0; i < le; i++ {
		fmt.Println("va", a[i])
	}
}

//range 循环可用于遍历数组中的元素。
func f4() {
	var a = [...]float64{1.2, 2.1, 3.4}
	for _, v := range a {
		fmt.Println("va", v)
	}
}

//多维数组
func f5() {
	var a = [3][2]string{
		{"w1", "w2"},
		{"x1", "x2"},
		{"s1", "s2"},
	}

	for i, v1 := range a {
		for x, v2 := range v1 {
			fmt.Printf("a[%d][%d]=%s ", i, x, v2)
		}
		fmt.Print("\r\n")
	}

	var b [3][2]string

	b[0][0] = "o_o_00"
	b[0][1] = "o_o_01"
	b[1][0] = "o_o_10"
	b[1][1] = "o_o_11"
	b[2][0] = "o_o_20"
	b[2][1] = "o_o_21"

	fmt.Println(b)
}

//切片
//切片的修改
//切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
func f6() {
	a := [4]int{1, 2, 3, 4}

	var c []int = a[:2]

	fmt.Println(c)

	b := []string{"2"}

	fmt.Println(b)

	var x = [...]int{1, 2, 3, 4, 5, 6, 7}

	xq := x[:3]
	fmt.Println(x)
	for i, _ := range xq {
		xq[i]++
	}
	fmt.Println(x)

}

//切片的长度和容量
//切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
func f7() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
}

// 创建切片
func f8() {
	v := make([]int, 7, 7)
	fmt.Println(v)

	s := []string{"s"}

	s = append(s, "99")

	fmt.Println(s)

	var w []int

	if w == nil {
		fmt.Println("88")
	}
	w = append(w, 9, 10, 20)
	fmt.Println(w)
	是 := []string{"99"}
	是1 := []string{"998", "99", "fff"}
	是 = append(是, 是1...)

	fmt.Println(是)

	//切片的函数传递
	one := []int{7, 8, 9}
	fmt.Println(one)
	subjectOne(one)
	fmt.Println(one)

	//多维切片
	var name = [][]string{
		{"1"},
		{"3"},
	}

	fmt.Println(name)

	// 切片复制

	f := []string{"s", "d", "s1", "d2", "s4", "d4", "s14"}
	fd := make([]string, len(f))

	copy(fd, f)
	fmt.Println(fd)
}

func subjectOne(num []int) {
	for i := range num {
		num[i] -= 2
	}
}
