package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var sum int64 = 0

	re := regexp.MustCompile(`\d`)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := re.FindAllString(line, -1)
		resu, _ := strconv.ParseInt(numbers[0]+numbers[len(numbers)-1], 10, 64)
		sum += resu
	}

	fmt.Println(sum)

	readFile.Close()
}
