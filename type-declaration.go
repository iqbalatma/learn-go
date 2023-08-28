package main

import "fmt"

func main() {
	type NoKTP string
	type Maried bool

	var noKtpIqbal NoKTP = "610107160299"
	var isMaried Maried = false

	fmt.Println(noKtpIqbal)
	fmt.Println(isMaried)
}
