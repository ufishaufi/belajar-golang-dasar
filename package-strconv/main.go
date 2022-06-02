package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	x := 2000000
	valueString := strconv.Itoa(x)
	// fmt.Println(value)
	fmt.Printf("'%v' is of type %T \n", valueString, valueString)


	valueInt, _ := strconv.Atoi("2000000")
	// fmt.Println(valueInt)
	fmt.Printf("%v is of type %T \n", valueInt, valueInt)
}