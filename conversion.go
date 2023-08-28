package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name string = "iqbal"
	var i byte = name[0]
	var iString string = string(i)

	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(iString)
}
