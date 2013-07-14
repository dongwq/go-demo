/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-11，上午11:15
 * @version 0.1
 */
package main

import (
	"jsondemo"
	"fmt"
)

func main() {

	jsondemo.MakeConfig()
	config := jsondemo.LoadConfig()

	fmt.Println(config)

}

