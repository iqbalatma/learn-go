package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}
func main() {
	fmt.Println(Ups(2))
}
