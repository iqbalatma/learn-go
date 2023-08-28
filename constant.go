package main

import "fmt"

func main() {
	const firstName string = "iqbal"
	const lastName = "atma"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		city     = "sambas"
		province = "kalimantan barat"
	)
}
