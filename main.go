package main

import (
	"fmt"
	"gomonitor/src/file"
	"gomonitor/src/monitor"
	"os"
)

func main() {
	showMenu()
	command := getAndValidateInput()
	handleOptions(command)
}

func showMenu() {
	fmt.Println("1. Start monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("3. Exit")
}

func getAndValidateInput() int {
	var command int
	_, err := fmt.Scan(&command)

	if err != nil || command < 1 || command > 3 {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	return command
}

func handleOptions(command int) {
	switch command {
	case 1:
		fmt.Println("")
		monitor.InitMonitor()
	case 2:
		fmt.Println("")
		file.ReadLogs()
		main()
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Something wrong happened.")
		os.Exit(-1)
	}
}
