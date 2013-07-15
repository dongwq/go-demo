package main

import (
	"fmt"
)

type TZ int

func (a *TZ) Incr(offset int) {
	*a += TZ(offset)
}

func main() {
	var i TZ

	i.Incr(100)

	fmt.Println(i)
}
