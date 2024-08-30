package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Read(path string) []string {
	fmt.Println("Getting file", path, "...")
	var sources []string

	filepath := getPath(path)
	file := createOrOpenFile(filepath, path)

	reader := bufio.NewReader(file)
	for {
		line, lineErr := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sources = append(sources, line)
		if lineErr == io.EOF {
			break
		}
	}

	defer file.Close()
	return sources
}

func ReadLogs() {
	fmt.Println("Getting file logs...")

	filepath := getPath(".log")
	file := createOrOpenFile(filepath, ".log")
	content, _ := io.ReadAll(file)

	fmt.Println(string(content))
}

func Write(path, content string) {
	filepath := getPath(path)
	file := createOrOpenFile(filepath, path)

	file.WriteString("[" + time.Now().Format("02/01/2006 15:04:05") + "]: " + content + "\n")

	defer file.Close()
}

func getPath(path string) (filePath string) {
	workingDirectory, wdErr := os.Getwd()
	if wdErr != nil {
		fmt.Println("Something wrong happened fetching working directory.", wdErr)
		os.Exit(1)
	}

	basePath := fmt.Sprintf("%s/%s", workingDirectory, path)
	return basePath
}

func createOrOpenFile(filepath, path string) *os.File {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Something wrong happened opening:", path, "error: ", err)
		os.Exit(1)
	}

	return file
}
