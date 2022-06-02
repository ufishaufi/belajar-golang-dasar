package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man Man) Marrieds() {
	man.Name = "Mr. " + man.Name
}

func main() {
	shaufi := Man{"Shaufi"}
	shaufi.Married()

	fmt.Println(shaufi.Name)

	shaufi_new := Man{"Shaufi"}
	shaufi_new.Marrieds()

	fmt.Println(shaufi_new.Name)
}