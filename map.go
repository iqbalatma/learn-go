package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "iqbal",
		"address": "selakau",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	fmt.Println(book)
}
