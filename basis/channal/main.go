/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-7，下午9:22
 * @version 0.1
 */
package main


import (
	"fmt"
	"time"
)

func main() {
	//创建一个string类型的channel
	channel := make(chan string)

	//创建一个goroutine向channel里发一个字符串
	go func() { channel <- "hello" }()

	msg := <- channel
	fmt.Println(msg)

}

func showMsg(chan string){
//	msg2 := <- chan
	// fmt.Println(msg2)
}



// chan as a queue, msg queue or ?? 
func main() {
	queue := make(chan int, 100)

	queue <- 1
	queue <- 2
	queue <- 3
	close(queue)
	go func() {
		fmt.Println("dequeue goroutine start------------------>")
__loop:
		for {
			select {
			case item, ok := <-queue:
				fmt.Println(item, ok)
				if !ok {
					break __loop
				}

			}
		}
		fmt.Println("dequeue goroutine exit------------------>")
	}()

	time.Sleep(1e18)

}

