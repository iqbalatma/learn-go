package main

import "fmt"

func sumAll(numbers ...int) int {
	var total int = 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	var total int = sumAll(1, 2, 3)

	fmt.Println(total)

	var slice = []int{
		1, 2, 3, 4, 5,
	}
	var total2 int = sumAll(slice...) //destruct slice to set into varargs

	fmt.Println(total2)

}
