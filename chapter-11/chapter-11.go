package main

import "fmt"

func main() {
	fmt.Println("MAP")

	var sl map[string]int
	if sl == nil {
		fmt.Println("map is nil. Going to make one.")
		sl = make(map[string]int)
	}
	sl["steve"] = 12000
	sl["jamie"] = 15000
	sl["mike"] = 9000

	slk := map[string]int{"sss": 1000, "xxxx": 200}

	v, ok := slk["sss"]
	if ok == true {
		fmt.Println(v)
	} else {
		fmt.Println(sl)
	}
	fmt.Println(sl)
	//有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。
	for i, val := range sl {
		fmt.Printf("k: %s v: %v", i, val)
	}
	delete(sl, "mike")
	fmt.Println(sl)
	fmt.Println("length is", len(sl))
	//和 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。

	//map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
}
