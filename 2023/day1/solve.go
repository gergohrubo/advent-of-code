package main

import (
	"aoc/2023/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/dlclark/regexp2"
)

func partOne() {
	var sum int64 = 0
	reg := regexp.MustCompile(`\d`)

	analyzeLine := func(line string) {
		numbers := reg.FindAllString(line, -1)
		res, _ := strconv.ParseInt(numbers[0]+numbers[len(numbers)-1], 10, 64)
		sum += res
	}

	utils.ReadFileByLine("input.txt", analyzeLine)

	fmt.Printf("Part 1: %v\n", sum)
}

func partTwo() {
	var sum int64 = 0
	reg := regexp2.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`, 0)
	// reg := regexp2.MustCompile(`(?=(\d|one|two|three|four|five|six|seven|eight|nine))`, 0)

	matchMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	getDigit := func(key string) string {
		result, found := matchMap[key]
		if !found {
			return key
		}
		return result
	}

	analyzeLine := func(line string) {
		match, _ := reg.FindStringMatch(line)
		firstMatch := match.String()
		lastMatch := match.String()

		for match != nil {
			lastMatch = match.String()
			match, _ = reg.FindNextMatch(match)
		}

		res, error := strconv.ParseInt(getDigit(firstMatch)+getDigit(lastMatch), 10, 64)
		if error != nil {
			log.Fatal("Got invalid number")
			log.Fatal(error)
		}
		sum += res
	}

	utils.ReadFileByLine("input.txt", analyzeLine)

	fmt.Printf("Part 2: %v\n", sum)
}

func main() {

	partOne()

	partTwo()
}
