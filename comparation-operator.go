package main

import "fmt"

func main() {
	var name1 string = "iqbal"
	var name2 string = "iqbal"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 int = 100
	var value2 int = 200

	fmt.Println(value2 > value1)
	fmt.Println(value2 < value1)
	fmt.Println(value2 == value1)
	fmt.Println(value2 != value1)
}
