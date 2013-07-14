/**
 * Created with IntelliJ IDEA.
 * User: dong1
 * Date: 13-3-25
 * Time: 下午2:40
 * To change this template use File | Settings | File Templates.
 */
package basic

import (
	"fmt"
	"time"
)

func DemoTicker() {

	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	//设置一个timer，10钞后停掉ticker
	timer := time.NewTimer(10*time.Second)
	<-timer.C

	ticker.Stop()
	fmt.Println("timer expired!")
}

