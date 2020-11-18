package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")

	app := "https://random-status-code.herokuapp.com"
	res, _ := http.Get(app)

	if res.StatusCode == 200 {
		fmt.Println("Application:", app, "was successfully loaded!")
	} else {
		fmt.Println("Application:", app, "is having problems. Status code:", res.StatusCode)
	}
}
