package main

import (
	"fmt"
	"time"
)

/**
并发是什么？
并发是指立即处理多个任务的能力。一个例子就能很好地说明这一点。

我们可以想象一个人正在跑步。假如在他晨跑时，鞋带突然松了。于是他停下来，系一下鞋带，接下来继续跑。这个例子就是典型的并发。这个人能够一下搞定跑步和系鞋带两件事，即立即处理多个任务。

并行是什么？并行和并发有何区别？
并行是指同时处理多个任务。这听起来和并发差不多，但其实完全不同。

我们同样用这个跑步的例子来帮助理解。假如这个人在慢跑时，还在用他的 iPod 听着音乐。在这里，他是在跑步的同时听音乐，也就是同时处理多个任务。这称之为并行。

Go 编程语言原生支持并发。
Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发。
*/
func main() {
	//Go 协程是什么？
	//Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。
	//与线程相比，创建一个 Go 协程的成本很小。因此在 Go 应用中，常常会看到有数以千计的 Go 协程并发地运行。

	go hello()
	time.Sleep(1 * time.Second)

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Go 协程")

}
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func hello() {
	fmt.Println("hello word!")
}
