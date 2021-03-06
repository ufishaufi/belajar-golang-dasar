package main

import "fmt"

type filter func(string) string

func sayHelloWithFilter(name string, filter filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ",  nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}


func main() {
	sayHelloWithFilter("Shaufi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}