package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 4
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
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
	fmt.Println("")

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	apps := readApps()

	for i := 0; i < monitoring; i++ {
		for i, app := range apps {
			fmt.Println("Testing application", i, ":", app)
			testApp(app)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testApp(app string) {
	res, err := http.Get(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if res.StatusCode == 200 {
		fmt.Println("Application:", app, "was successfully loaded!")
		registryLog(app, true)
	} else {
		fmt.Println("Application:", app, "is having problems. Status code:", res.StatusCode)
		registryLog(app, false)
	}
}

func readApps() []string {
	var apps []string

	file, err := os.Open("apps.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)
		apps = append(apps, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return apps
}

func registryLog(app string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + app + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(file))
}
