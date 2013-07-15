package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) add() bool {
	fmt.Println("user add")
	return false
}

func main() {
	fmt.Println("Hello")

	u := User{Name: "ok", Age: 2}

	fmt.Println(u)

}
