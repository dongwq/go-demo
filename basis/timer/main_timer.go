/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-7，下午9:46
 * @version 0.1
 */
package main

import (
	"time"
	"fmt"
	"timer"
)

func TimeSleep(){
	timer2 := time.NewTimer( 2 * time.Second)

	go func(){
		//		for{
		//			fmt.Println("OK")
		//		}

	}()
	<- timer2.C

	fmt.Println("timer run over")
}

func main(){
	timer.TimerEvent()
}

