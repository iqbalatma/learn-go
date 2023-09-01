package main

import "fmt"

func main() {
	var name string = "iqbal"
	if name == "iqbal" {
		fmt.Println(name)
	}

	var lastname string = "atma muliawan"

	if lastname == "atma" {
		fmt.Println(lastname)
	}

	//	short if statement
	if length := len(name); length > 10 {
		fmt.Println("terlalu panjang")
	}
}
