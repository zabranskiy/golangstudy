package main

import (
	"fmt"
)

func shift(s []int, n int) {
	c := len(s)
	reverse(s, n)
	fmt.Println(s)
	reverse(s[n:], c - n)
	fmt.Println(s)
	reverse(s, c)
}
func reverse(s []int, d int) {
	for i := 0; i < (d / 2); i++ {
		s[i], s[d-i-1] = s[d-i-1], s[i]
	}
}

func main() {
	myArray := [4]int{1, 2, 3, 4}
	var s []int = myArray[:]
	fmt.Println(s)
	shift(s, 2)
	fmt.Println(s)
	myArray[2] = 5
	fmt.Println(s)
	reverse(s, len(s))
	fmt.Println(s)
	myArray[3] = 7
	fmt.Println(s)
}
