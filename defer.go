package main

import "fmt"

// defer is a function that always execute even error is happened
// why we need to call logging with defer ? to make sure logging always execute
// when we not use defer and put logging in last line of code, when error happen logging will not be executed
func logging() {
	fmt.Println("Aplikasi selesai di jalankan")
}
func main() {
	defer logging()
}
