package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
	// Married		  bool	
}

func main() {
	var shaufi Customer
	shaufi.Name = "Shaufi"
	shaufi.Address = "Indonesia"
	shaufi.Age =29

	fmt.Println(shaufi.Name)
	fmt.Println(shaufi.Address)
	fmt.Println(shaufi.Age)

	joko := Customer {
		Name: 		"Joko",
		Address: 	"Cirebon",
		Age:		35,
	}
	fmt.Println(joko)
	
	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
}