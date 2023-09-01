package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "iqbal"
	names[1] = "atma"
	names[2] = "muliawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var sample [10]string

	fmt.Println(len(values))
	fmt.Println(len(sample))
}
