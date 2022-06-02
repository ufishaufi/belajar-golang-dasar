package main

import "fmt"

type Customer struct {
	Name, Address string
	Age			  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name) 
}

func (a Customer) sayHuu() {
	fmt.Println("Huuuuu from", a.Name)
}

func main() {
	var shaufi Customer
	shaufi.Name = "Shaufi"
	shaufi.Address = "Indonesia"
	shaufi.Age = 29

	shaufi.sayHello("Joko")
	shaufi.sayHuu()
}