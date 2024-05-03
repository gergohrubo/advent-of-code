package main

import (
	"aoc/2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Numbermap struct {
	num   int64
	index []int
}

func checkCurrentNumber(numberIndexPair []int, symbolIndexes [][]int) bool {
	belowThreshold := numberIndexPair[0] - 1
	aboveThreshold := numberIndexPair[1]
	for _, symbolIndexPair := range symbolIndexes {
		if symbolIndexPair[0] >= belowThreshold && symbolIndexPair[0] <= aboveThreshold {
			return true
		}
	}
	return false
}

func main() {
	var sum int64 = 0
	regSymbols := regexp.MustCompile(`[^\d|\.]`)
	regNumbers := regexp.MustCompile(`([1-9][0-9]*)`)
	previousSymbolIndexes := make([][]int, 0)
	previousNumbers := make([]Numbermap, 0)

	analyzeLine := func(line string) {
		fmt.Printf("Line: %v\n", line)
		currentSymbolIndexes := regSymbols.FindAllStringIndex(line, -1)
		numberIndexes := regNumbers.FindAllStringIndex(line, -1)

		for _, numbermap := range previousNumbers {
			fmt.Printf("Previous number element: %v\n", numbermap.num)
			found := checkCurrentNumber(numbermap.index, currentSymbolIndexes)
			if found {
				fmt.Printf("Adding previous number: %v\n", line[numbermap.index[0]:numbermap.index[1]])
				sum += numbermap.num
			}
		}

		previousNumbers = nil
		symbolIndexes := append(previousSymbolIndexes, currentSymbolIndexes...)

		for _, numberIndexPair := range numberIndexes {
			number, _ := strconv.ParseInt(line[numberIndexPair[0]:numberIndexPair[1]], 10, 64)
			found := checkCurrentNumber(numberIndexPair, symbolIndexes)
			if found {
				fmt.Printf("Adding number: %v\n", line[numberIndexPair[0]:numberIndexPair[1]])
				sum += number
			} else {
				fmt.Printf("Adding this to previous number slice: %v\n", line[numberIndexPair[0]:numberIndexPair[1]])
				previousNumbers = append(previousNumbers, Numbermap{num: number, index: numberIndexPair})
			}
		}

		previousSymbolIndexes = currentSymbolIndexes
	}
	utils.ReadFileByLine("input.txt", analyzeLine)

	fmt.Printf("Sum: %v\n", sum)

	// re := regexp.MustCompile(`\d`)
	// re3 := regexp.MustCompile(`([1-9][0-9]*)`)
	// re2 := regexp.MustCompile(`[^\d|\.]`)
	// fmt.Println(re.FindStringIndex("t1able345tt"))
	// fmt.Println(re.FindAllStringIndex("t1able345tt", -1))
	// fmt.Println(re3.FindAllStringIndex("t1able345tt", -1))
	// fmt.Println(re3.FindAllStringSubmatchIndex("t1able345tt", -1))
	// fmt.Println(re2.FindAllStringIndex("t1ab.l..e345tt", -1))
	// fmt.Println(re.FindStringIndex("foo") == nil)
}
