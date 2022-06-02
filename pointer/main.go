package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func pass_by_value() {

	// pass by value
	alamat1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	alamat2 := alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1) // alamat1 tidak berubah
	fmt.Println(alamat2)
}


func pass_by_reference() {
	
	// Pointer (Pass by reference) with only operator &
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	// tidak berubah isi data address1 dan address3 kalau hanya menggunakan operator &
	address2 = &Address{"Malang", "Jawa Timus", "Indonesia"}
	
	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)
}

func pointer_with_operator_asterisk() {

	// Pointer (Pass by reference) with only operator & and *
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"
	
	// merubah isi data address1 dan address3 karena dia merefer ke memori Address
	*address2 = Address{"Malang", "Jawa Timus", "Indonesia"}

	fmt.Println(address1) // address 1 berubah
	fmt.Println(address2)
	fmt.Println(address3)
}

func pointer_with_new() {

	// new function
	var address4 *Address = new(Address)

	address4.City = "Jakarta"

	fmt.Println(address4)

}

func main() {

	pass_by_value()

	fmt.Println()

	pass_by_reference()

	fmt.Println()

	pointer_with_operator_asterisk()

	fmt.Println()

	pointer_with_new()

}