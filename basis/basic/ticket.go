/**
 * Created with IntelliJ IDEA.
 * User: dong1
 * Date: 13-3-25
 * Time: 上午11:19
 * To change this template use File | Settings | File Templates.
 */
package basic

import (
	"time"
	"math/rand"
	"fmt"
	"runtime"
	"sync"
)


var total_tickets int32 = 10;
var mutex sync.Mutex

func sell_tickets(i int){
	for total_tickets > 0{

		mutex.Lock()
		if total_tickets > 0 { //如果有票就卖
			time.Sleep( time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets-- //卖一张票
			fmt.Println("id:", i, "  ticket:", total_tickets)
		}else{
			break
		}
		mutex.Unlock()
	}
}

func DemoTicket() {
	runtime.GOMAXPROCS(4) //我的电脑是4核处理器，所以我设置了4
	rand.Seed(time.Now().Unix()) //生成随机种子

	for i := 0; i < 5; i++ { //并发5个goroutine来卖票
		go sell_tickets(i)
	}
	//等待线程执行完
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done") //退出时打印还有多少票
}
