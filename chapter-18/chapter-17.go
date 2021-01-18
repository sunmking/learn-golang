package main

import (
	"fmt"
	"time"
)

/**
什么是信道？
信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收。
*/
func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T \r\n", a)
	}
	go hello(a)
	<-a
	fmt.Println("Go 信道")
	// 其他实例
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
	// 死锁
	// fatal error: all goroutines are asleep - deadlock!
	//xc := make(chan int)
	//xc <- 6
	// 单信道
	sn := make(chan<- int)
	go sendData(sn)
	//fmt.Println(<-sn)
}
func sendData(sendch chan<- int) {
	sendch <- 10
}
func hello(done chan int) {
	time.Sleep(4 * time.Second)
	fmt.Println("hello word")
	done <- 1
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
