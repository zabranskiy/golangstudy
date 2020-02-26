package main

import (
	"fmt"
	"reflect"
)

func cyclicShift(s interface{}, n int) (out []interface{}) {
	sVal := reflect.ValueOf(s)
	if sVal.Kind() == reflect.Slice {
		c := sVal.Len()
		out = make([]interface{}, c)
		k := 0
		for i := c-n; i < c; i, k = i+1, k+1 {
			out[k] = sVal.Index(i).Interface()
		}
		for i := 0; i < c-n; i, k = i+1, k+1 {
			out[k] = sVal.Index(i).Interface()
		}
		return out
	} else {
		panic(sVal.Kind())
	}
}

func revertSlice(s interface{}) (out []interface{}) {
	sVal := reflect.ValueOf(s)
	if sVal.Kind() == reflect.Slice {
		c := sVal.Len()
		out = make([]interface{}, c)
		for i := 0; i < c; i++ {
			out[c-i-1] = sVal.Index(i).Interface()
		}
		return out
	} else {
		panic(sVal.Kind())
	}
}

func main() {
	myArray := [4]int{1, 2, 3, 4}
	myArray2 := [5]string{"a", "b", "c", "d", "e"}
	var s []int = myArray[:]
	var s2 []string = myArray2[:]
	fmt.Println(s)
	fmt.Println(s2)
	res := cyclicShift(s, 1)
	fmt.Println(res)
	myArray2[2] = "w"
	fmt.Println(s2)
	fmt.Println(res)
	res2 := revertSlice(s)
	fmt.Println(res2)

}
