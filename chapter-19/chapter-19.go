package main

import (
	"fmt"
	"time"
)

//无缓冲信道的发送和接收过程是阻塞的。
//我们还可以创建一个有缓冲（Buffer）的信道。
//只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
//同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
func main() {
	fmt.Println("缓冲信道")

	lch := make(chan string, 2)
	lch <- "hello"
	lch <- "word"
	fmt.Println(<-lch)
	fmt.Println(<-lch)

	done := make(chan int, 2)
	go write(done)
	time.Sleep(2 * time.Second)
	for v := range done {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}

	// 死锁  缓冲区满或者为空 都会造成死锁
	don := make(chan int, 2)
	don <- 1
	don <- 2
	//don<-3
	fmt.Println(<-don)
	fmt.Println(<-don)
	//fmt.Println(<-don)

	// 长度 & 容量
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
