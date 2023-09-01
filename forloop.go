package main

import "fmt"

func main() {
	var counter int = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 20; counter2++ {
		fmt.Println("Looping2 Ke ", counter2)
	}

	var slice = []string{
		"iqbal",
		"atma",
		"muliawan",
	}
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, item := range slice {
		fmt.Println("Index ", index, "=", item)
	}

	for _, item := range slice {
		fmt.Println("value =", item)
	}
}
