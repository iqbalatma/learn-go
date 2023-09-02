package main

import "fmt"

func getFullNameWithMiddlename() (firstName string, middleName string, lastName string) {
	firstName = "iqbal"
	middleName = "atma"
	lastName = "muliawan"

	return
}

func main() {
	firstname, middlename, lastname := getFullNameWithMiddlename()

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}
