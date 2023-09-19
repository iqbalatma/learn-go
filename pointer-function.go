package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "indonesia"
}

func main() {
	var address1 = Address{City: "Skw", Province: "Kalbar", Country: "sgp"}
	changeAddressToIndonesia(&address1)

	fmt.Println(address1.Country)
}
