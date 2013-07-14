import (
  "fmt"
	"github.com/xiocode/rss"
	"io/ioutil"
	"net/http"
	// "runtime"
	"sync"
	"time"
)
 
type Feed *rss.Feed
 
type Subscriber struct {
	mutex         sync.Mutex
	wg            *sync.WaitGroup // Fetch wg
	in            chan string     // RSS URL
	out           chan Feed       // WebBody
	done          chan bool       // done
	doneWaitGroup *sync.WaitGroup
	ticker        *time.Ticker
	shuttingDown  bool // 正在关闭
	shutDown      chan bool
	limit         chan int // crawler limit
}
 
func (this *Subscriber) Init() {
	this.doneWaitGroup = new(sync.WaitGroup) // waiting shutdown
	this.wg = new(sync.WaitGroup)            // waiting crawl done
	this.in = make(chan string, 10)          // RSS URL
	this.done = make(chan bool)              // rss crawl done
	this.out = make(chan Feed, 10)           // WebBody
	this.ticker = time.NewTicker(10 * time.Second)
	this.limit = make(chan int, 100)
}
 
func (this *Subscriber) cralwer(url string) {
	this.limit <- 1
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	body = encoding(body)
	feed, err := rss.Parse(body)
	this.out <- feed
}
 
func encoding(content []byte) []byte {
	return content
}
 
// 分配工
func (this *Subscriber) mux() {
loop:
	for {
		select {
		case url := <-this.in:
			this.wg.Add(1)
			go func(url string) {
				fmt.Println("CRAWLING")
				this.cralwer(url)
				this.wg.Done()
				<-this.limit
				fmt.Println("DONE!")
			}(url)
		case feed := <-this.out: // TODO 处理RSS
			this.wg.Add(1)
			go func(feed Feed) {
				this.store(feed)
				this.wg.Done()
			}(feed)
		case _ = <-this.done: // shutdown 关闭
			go func() {
				for _ = range this.done {
				}
			}()
			break loop
		case _ = <-this.ticker.C: // 心跳检测
			fmt.Println("10 SECOND")
		}
	}
	this.wg.Wait()
}
 
func (this *Subscriber) store(feed Feed) {
	if feed == nil {
		return
	}
	Store(feed)
}
 
func (this *Subscriber) Subscribe(url string) {
	fmt.Println("添加新聞源地址: " + url)
	go func() {
		this.in <- url
		fmt.Println("添加新聞源成功！")
	}()
}
 
func (this *Subscriber) Start() (done *sync.WaitGroup) {
 
	this.doneWaitGroup.Add(1)
	go func() {
		this.mux()
		fmt.Println("结束！")
		this.doneWaitGroup.Done()
	}()
 
	done = this.doneWaitGroup
	return
}
 
func (this *Subscriber) Done() {
	this.done <- true
	this.shuttingDown = true
	this.doneWaitGroup.Wait()
}
 
func (this *Subscriber) Shutdown() {
	if this.shuttingDown {
		return
	}
	this.shuttingDown = true
	this.doneWaitGroup.Add(1)
	this.Done()
	this.doneWaitGroup.Done()
}