/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-11，下午4:32
 * @version 0.1
 */
package netdemo

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func ClientDemo() {

//	client := http.DefaultClient
	resp, err := http.Get("https://account.xiaomi.com/pass/serviceLogin");
	if err != nil{
		log.Fatal("get err", err)
	}

	fmt.Println(resp)

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(data))
}

