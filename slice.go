package main

import "fmt"

func main() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(months)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"

	fmt.Println(months)
	fmt.Println(slice1)

	slice1[1] = "Mei"

	fmt.Println("Diubah ke mei")
	fmt.Println(months)
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Iqbal")
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "iqbal"
	newSlice[1] = "atma"

}
