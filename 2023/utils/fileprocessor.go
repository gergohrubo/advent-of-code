package utils

import (
	"bufio"
	"fmt"
	"os"
)

type ReadFileCallback func(string)

func ReadFileByLine(path string, callback ReadFileCallback) {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		callback(line)
	}

	readFile.Close()
}
