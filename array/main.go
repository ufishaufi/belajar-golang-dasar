package main

import "fmt"

func main() {
	
	var names [3]string

	names[0] = "muhammad"
	names[1] = "shaufi"
	names[2] = "imanulhaq"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi[10]string

	fmt.Println(len(lagi))

	

}