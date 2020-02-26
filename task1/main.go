package main

import (
	"fmt"
)

func cyclicShift(s []int, n int) {
	c := len(s)
	tmp := make([]int, c)
	copy(tmp, s)
	k := 0
	for i := c-n; i < c; i, k = i+1, k+1 {
		s[k] = tmp[i]
	}
	for i := 0; i < c-n; i, k = i+1, k+1 {
		s[k] = tmp[i]
	}
}
func revertSlice(s []int) {
	c := len(s)
	for i := 0; i < (c / 2); i++ {
		tmp := s[i]
		s[i] = s[c-i-1]
		s[c-i-1] = tmp
	}
}

func main() {
	myArray := [4]int{1, 2, 3, 4}
	var s []int = myArray[:]
	fmt.Println(s)
	cyclicShift(s, 1)
	fmt.Println(s)
	myArray[2] = 5
	fmt.Println(s)
	revertSlice(s)
	fmt.Println(s)
	myArray[3] = 7
	fmt.Println(s)
}
