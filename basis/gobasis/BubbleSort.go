package main

import (
	"fmt"
)

func BubbleSort(a []int) {
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}

		}
	}

	//return a
}

func main() {
	// can't use ... in []
	a := []int{1, 23, 4, 5, 8, 33, 70, 3, 323}

	fmt.Println(a)

	BubbleSort(a)

	fmt.Println(a)

}
