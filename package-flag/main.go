package main

import (
	"fmt"
	"flag"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)
	fmt.Println("Number: ", *number)

	// go run main.go -host=localhost -user=root -password=root
	// go run main.go -host=localhost -user=root -password=root -number=90
}