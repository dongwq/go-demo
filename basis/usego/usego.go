/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-10，上午10:00
 * @version 0.1
 */
package usego

import (
	"fmt"
)

func CountTimes() {

	data := []int{2, 3, 5, 9, 3}

	countMap := make(map[int]int)

	for i := 0; i < len(data); i++ {
		d := data[i]
		fmt.Printf("%d,", d)
		if v, ok := countMap[d]; ok {
			countMap[d] = v + 1
		} else {
			countMap[d] = 1
		}
	}
	fmt.Println()
	fmt.Println(countMap)

}

func GetMin() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := 1000;
	//	for i := 0; i < len(x); i++ {
	//		xi := x[i]
	//		if (xi < min) {
	//			min = xi
	//		}
	//	}

	for i, v := range x {
		fmt.Printf("i=%d, v=%d\n", i,v)
		xi := v
		if (xi < min) {
			min = xi
		}
	}
	fmt.Printf("min=%d", min)

}
