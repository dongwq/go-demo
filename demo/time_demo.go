// time_demo
package demo

import (
	"fmt"
	"time"
)

func TimeDemo() {
	//Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub，将t0加d获取t1就是使用Add方法
	k := time.Now()
	//一天之前
	d, _ := time.ParseDuration("-24h")
	fmt.Println(k.Add(d))
	//一周之前
	fmt.Println(k.Add(d * 7))
	//一月之前
	fmt.Println(k.Add(d * 30))

	fmt.Println(k.Format(time.RFC3339))
	fmt.Println(k.Format("2006-01-02 15:04:05"))

	//2006-01-02T15:04:05Z07:00

}
