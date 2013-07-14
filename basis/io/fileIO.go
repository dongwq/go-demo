/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-7，下午4:32
 * @version 0.1
 */
package main

import (
	log "github.com/cihub/seelog"
	"fmt"
	"os"
	"container/list"
)

func main() {
	f, err := os.Open("ok.txt");
	if err != nil {
		log.Info(err)
		return;
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		log.Error("stat err")
	}
	data := make([]byte, stat.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Error("read err")
	}
	dataStr := string(data)

	log.Info(dataStr)


	dirInfo()

	demoList()

}

func dirInfo() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func demoList() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
