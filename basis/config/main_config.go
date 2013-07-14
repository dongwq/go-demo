/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-13，下午11:34
 * @version 0.1
 */
package main

import (
	"config"
	"log"
	"fmt"
	"crypto/aes"
)

func main() {
	conf, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal("load err", err)
	}

	fmt.Println(conf)
	var p = conf["Phone"]
	fmt.Printf("%v,%T\n",p, p)
	fmt.Println(conf["Phone"])

	k := []byte("efwfw")
	c, err := aes.NewCipher(k)
	if err == nil {
		log.Print("Hello")
		log.Print(err)
	}
	log.Print(c)
	//	var dst []byte = make([]byte,50)
	//	var src string = "ok";
	//	src.getBytes()
	//	c.Encrypt(dst,[]byte("ok") )
	//	log.Print(dst)
}

