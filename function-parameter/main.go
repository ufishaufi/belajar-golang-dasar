package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Muhammad"
	sayHelloTo(firstName, "Shaufi")
	sayHelloTo("Budi", "Nugraha")
}