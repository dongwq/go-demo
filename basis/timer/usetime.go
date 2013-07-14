/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-10，下午12:28
 * @version 0.1
 */
package timer

import (
	"time"
	"fmt"
)

func TimerEvent(){
	ticker := time.NewTicker(time.Second)

	go func () {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	//设置一个timer，10钞后停掉ticker
	timer := time.NewTimer(10*time.Second)
	<- timer.C

	ticker.Stop()
	fmt.Println("timer expired!")
}



