package main

import (
	"fmt"
	"unicode/utf8"
)

func CompareChars(word string) {
	s := []byte(word)
	for utf8.RuneCount(s) > 1 {
		r, size := utf8.DecodeRune(s)
		s = s[size:]
		nextR, size := utf8.DecodeRune(s)
		fmt.Print(r == nextR, ",")
	}
	fmt.Println()
}

func main() {

	fmt.Println(len("Hello, 世界")) // 13

	hello := "Hello, 世界"
	for i := range hello {
		fmt.Print(string(hello[i]))
	}
	
	fmt.Println()
	fmt.Println(utf8.RuneCountInString("Hello, 世界"))

	CompareChars("hello")
	
	CompareChars("你好好好")


}
