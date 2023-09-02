package main

import "fmt"

type Blacklist func(string) bool

func halloWithBlacklist(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Saya di block")
	} else {
		fmt.Println("Lanjuutt")
	}
}

func main() {
	halloWithBlacklist("anjing", func(name string) bool {
		return name == "anjing"
	})
}
