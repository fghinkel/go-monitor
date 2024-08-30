package monitor

import (
	"fmt"
	"gomonitor/src/file"
	"net/http"
	"os"
	"time"
)

const DEFAULT_DELAY = 5

func InitMonitor() {
	fmt.Println("Starting monitor...")
	sources := file.Read("sources.txt")
	i := 0

	for {
		time.Sleep(DEFAULT_DELAY * time.Second)
		i++

		fmt.Println("Iteration.", i)
		for _, source := range sources {
			handleRequest(source)
		}

		fmt.Println("") //empty print to break line on cli
	}
}

func handleRequest(source string) {
	response, err := http.Get(source)

	if err != nil {
		fmt.Println("Something wrong happened:", err)
		os.Exit(1)
	}

	if response.StatusCode == 200 {
		message := fmt.Sprintf("%s is running with status: %s", source, response.Status)

		fmt.Println(message)
		file.Write(".log", message)
	} else {
		message := fmt.Sprintf("Something wrong happened %s, returns status: %s", source, response.Status)

		file.Write(".log", message)
		fmt.Println(message)
	}
}
