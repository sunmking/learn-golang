package main

import "fmt"

func main() {
	fmt.Println("可变参数函数")
	num := 1
	nums := []int{1, 2, 3, 4, 5}
	find(num, nums...)

	chang := []string{"Hello", "World"}
	change(chang...)
	fmt.Println(chang)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if num == v {
			fmt.Printf("%d,%v", i, v)
			found = true
		}
	}

	if !found {
		fmt.Println("no found")
	}
}

func change(s ...string) {
	s[0] = "Go"
}
