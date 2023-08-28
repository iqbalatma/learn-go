package main

import "fmt"

func main() {
	var name string

	//data type initiate on definition
	name = "Iqbal Atma muliawan"
	fmt.Println(name)

	//modify variable
	name = "atma dev"
	fmt.Println(name)

	//data type based on initial value
	var address = "selakau"
	fmt.Println(address)

	//declaration without var
	kabupaten := "Sambas"
	fmt.Println(kabupaten)

	//multiple declaration
	var (
		firstName = "iqbal"
		lastname  = "atma"
	)

	fmt.Println(firstName)
	fmt.Println(lastname)
}
