// DemoOne
package main

import (
	"fmt"
)

//a give a PingPong by chan

var c chan string

func PingPong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From PingPong:%d", i)
		i++

	}
}

func main() {
	c = make(chan string)

	go PingPong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main:%d", i)
		fmt.Println(<-c)
	}

	c2 := make(chan bool)
	go func() {
		fmt.Println("this a c2")
		c2 <- true
	}()

	<-c2
}



