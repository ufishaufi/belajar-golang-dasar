package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPShaufi NoKTP = "18741982741897419874"
	var marriedStatus Married = false
	fmt.Println(noKTPShaufi)
	fmt.Println(marriedStatus)
}