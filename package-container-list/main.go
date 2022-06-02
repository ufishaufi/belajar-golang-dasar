package main


import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()
	
	data.PushBack("Muhammad")
	data.PushBack("Shaufi")
	data.PushBack("Imanulhaq")
	data.PushFront("Budi")

	// fmt.Println(data.Front().Next())
	// fmt.Println(data.Front().Prev())

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println()

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}