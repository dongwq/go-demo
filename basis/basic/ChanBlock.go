package basic

import (
	"fmt"
	"time"
)

func ChanBlock333() {

	channel := make(chan string) //注意: buffer为1

	go func() {
		channel <- "hello"
		fmt.Println("write \"hello\" done!")

		channel <- "World" //Reader在Sleep，这里在阻塞
		fmt.Println("write \"World\" done!")

		fmt.Println("Write go sleep...")
		time.Sleep(3*time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2*time.Second)
	fmt.Println("Reader Wake up...")

	msg := <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel //Writer在Sleep，这里在阻塞
	fmt.Println("Reader: ", msg)
}

// change reader block??
func ChanBlock2() {

	var ch1 chan string = make(chan string)


	go func() {
		for {
			v2 := <-ch1
			fmt.Println("Reader: ", v2)
		}
	}()

	var v string
	for {
		fmt.Scanln(&v)
		fmt.Println("Reader in: ", v)
		ch1 <- v
		//		time.Sleep(2*time.Second)

	}
}


func ChanSelect() {
	//创建两个channel - c1 c2
	c1 := make(chan string)
	c2 := make(chan string)

	//创建两个goruntine来分别向这两个channel发送数据
	go func() {
		time.Sleep(time.Second*1)
		c1 <- "Hello"
	}()
	go func() {
		time.Sleep(time.Second*1)
		c2 <- "World"
	}()

	//使用select来侦听两个channel
	////	for i := 0; i < 2; i++ {
	//		select {
	//		case msg1 := <-c1:
	//			fmt.Println("received", msg1)
	//		case msg2 := <-c2:
	//			fmt.Println("received", msg2)
	//		}
	////	}

	//	for {
	timeout_cnt := 0
	select {
	case msg1 := <-c1:
		fmt.Println("msg1 received", msg1)
	case msg2 := <-c2:
		fmt.Println("msg2 received", msg2)
	case <- time.After(time.Second*30) :
		fmt.Println("Time Out")
		timeout_cnt++
	}
	//		if timeout_cnt > 3 {
	//			break
	//		}
	//	}
}

