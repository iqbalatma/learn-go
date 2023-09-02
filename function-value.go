package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	var goodbye = getGoodBye

	fmt.Println(goodbye("iqbal"))
}
