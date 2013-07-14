// godemo project main.go
package main

/**
这个程序看到几点：

1 TypeOf和ValueOf是获取Type和Value的方法
2 ValueOf返回的<float64 Value>是为了说明这里的value是float64
3 第三个b的定义实现了php中的string->method的方法，为什么返回的是reflect.Value[]数组呢？当然是因为Go的函数可以返回多个值的原因了。

*/
import (
	"fmt"
	 "github.com/dongwq/godemo/demo"
	"github.com/dongwq/godemo/ref"
	_ "github.com/dongwq/godemo/util"
)

func TestRef3() {
	// iterate through the attributes of a Data Model instance
	for name, mtype := range ref.Attributes(&ref.Dish{}) {
		fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	}

}



func main() {
	//demo.Ref1()
	//demo.TestRef3()
	//demo.DemoMysql();
	//demo.DemoRegexp();
	demo.DemoPinyin();

	//demo.Md5Demo();

	//demo.TimeDemo()
	//demo.DemoAES()
	//demo.DemoAESStr()
	
	//demo.DownImage()
	//demo.DemoHijack();
	//util.DemoMail()
	//_ = ref.RefAction
	//ref.RefAction()
	//ref.DemoMapFunc();
}
