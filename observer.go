package main

import "fmt"

func main() {
	username := "Alan"
	version := 1.0

	fmt.Println("Hello, sr.", username)
	fmt.Println("This application is at version:", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exting")
	default:
		fmt.Println("Unknown command")
	}
}
