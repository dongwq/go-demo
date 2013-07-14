/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-7，下午12:36
 * @version 0.1
 */
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b
}

func mmsg() (int, string, int) {
	return 1, "msg", 2
}

func sum(nums ...int) {
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func demoMulti() {
	fmt.Println(max(4, 5))

	var a, b, c = mmsg()
	fmt.Printf("a=%d,b=%s,c=%d\n", a, b, c)

	// demo ...
	sum(1, 2)
	sum(1, 2, 3)

	//传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
func main() {
	// demo fact
	fmt.Println(fact(7))

	// demo defer
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	var d complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", d)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	//err :=
	err2 := json.Unmarshal(jsonBlob, &animals)
	if err2 != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	fmt.Printf("%t\n", 1 == 2)

	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("八进制：%o\n", 255)
	fmt.Printf("十六进制：%X\n", 255)
	fmt.Printf("十进制：%d\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("字符串：%s\n", "hello world")

}
