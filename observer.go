package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(-1)
	}
}

func showIntro() {
	username := "Alan"
	version := 1.0

	fmt.Println("Hello, sr.", username)
	fmt.Println("This application is at version:", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}
