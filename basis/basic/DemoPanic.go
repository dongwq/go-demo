/**
 * Created with IntelliJ IDEA.
 * User: dong1
 * Date: 13-3-25
 * Time: 上午11:02
 * To change this template use File | Settings | File Templates.
 */
package basic

import "fmt"

func g(i int) {
	if i > 1 {
		fmt.Println("Panic!")
		panic(fmt.Sprintf("%v", i))
	}

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Calling g with ", i)
		g(i)
		fmt.Println("Returned normally from g.")
	}
}

func DemoPanic() {
	f()
	fmt.Println("Returned normally from f.")
}

