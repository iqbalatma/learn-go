package main

import "fmt"

// by default golang using pass by value not by reference
// pointer is ability to reference into memory, not duplicate the data
type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"bandung", "jabar", "indonesia"}
	var address2 *Address = &address1

	address2.Country = "Singapore"
	fmt.Println(address1)
	fmt.Println(address2)

	var address3 Address = Address{"bandung", "jabar", "indonesia"}
	var address4 *Address = &address1

	address3.Country = "Singapore"

	*address4 = Address{"selakau", "kalbar", "indo"}
	fmt.Println(address3)
	fmt.Println(address4)

	var address5 *Address = new(Address)
	fmt.Println(address5)
}
