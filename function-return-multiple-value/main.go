package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhammad", "Shaufi", "Imanulhaq"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

}