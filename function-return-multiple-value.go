package main

import "fmt"

func getFullName() (string, string) {
	return "iqbal", "atma"
}
func main() {
	firstName, lastName := getFullName()
	firstName1, _ := getFullName()

	fmt.Println(firstName1)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
