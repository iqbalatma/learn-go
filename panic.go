package main

import "fmt"

// panic function use to stop program execution
// defer still execute even we found panic

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("pesan error ", message)
	}
	fmt.Println("applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("aplikasi berjalan")

}
func main() {
	runApp(true)
}
