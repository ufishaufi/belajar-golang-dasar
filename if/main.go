package main

import "fmt"

func main() {
	var name = "Kurniawan"

	if name == "Shaufi" {
		fmt.Println("Hello Shaufi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Boleh Kenalan")
	}

	if length := len(name) ;length > 5 { 
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}