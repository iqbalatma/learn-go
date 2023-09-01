package main

import "fmt"

func main() {
	var name = "iqbal"

	switch name {
	case "iqbal":
		fmt.Println("helloworld")
		fmt.Println("helloworld")
	case "atma":
		fmt.Println("Hello dunia")
		fmt.Println("Hello dunia")
	default:
		fmt.Println("Hello dunia satu")
	}

	//	switch dengan short statement
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("helloworld")
		fmt.Println("helloworld")
	case false:
		fmt.Println("Hello dunia")
		fmt.Println("Hello dunia")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("helloworld")
		fmt.Println("helloworld")
	case length >= 10:
		fmt.Println("Hello dunia")
		fmt.Println("Hello dunia")
	default:
		fmt.Println("Hello dunia")
	}
}
