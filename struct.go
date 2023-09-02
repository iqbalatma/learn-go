package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hallo ", customer.Name)
}

func main() {
	var iqbal Customer
	iqbal.Age = 25
	iqbal.Name = "atma"
	iqbal.Address = "selakau"

	fmt.Println(iqbal)

	atma := Customer{
		Name:    "atma",
		Address: "selakau",
		Age:     1,
	}

	fmt.Println(atma)

	atma.sayHi()
}
