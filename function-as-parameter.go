package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	var nameFiltered = filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func sayHelloWithFilter2(name string, filter Filter) { //if function declaration is too long, we can change it into type declaration
	var nameFiltered = filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("anjing", spamFilter)
}
