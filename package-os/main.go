package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :", args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	// go run main.go muhammad shaufi imanulhaq
	// export APP_USERNAME=root
	// export APP_PASSWORD=root
}