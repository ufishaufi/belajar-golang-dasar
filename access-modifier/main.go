package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Shaufi")
	// helper.sayGoodBye("Shaufi") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}