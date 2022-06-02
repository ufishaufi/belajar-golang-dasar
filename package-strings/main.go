package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Shaufi", "Shaufi"))
	fmt.Println(strings.Contains("Muhammad Shaufi", "Budi"))

	fmt.Println(strings.Split("Muhammad Shaufi Imanulhaq", " "))

	fmt.Println(strings.ToLower("Muhammad Shaufi Imanulhaq"))
	fmt.Println(strings.ToUpper("Muhammad Shaufi Imanulhaq"))
	fmt.Println(strings.ToTitle("Muhammad Shaufi Imanulhaq"))

	fmt.Println(strings.Trim("     Shaufi Imanulhaq      ", " "))

	fmt.Println(strings.ReplaceAll("Shaufi Joko Budi", "Shaufi", "Budi"))
	
}