package main

import (
	"aoc/2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

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

func partOne() {
	var sum int64 = 0
	regSymbols := regexp.MustCompile(`[^\d|\.]`)
	regNumbers := regexp.MustCompile(`([1-9][0-9]*)`)
	previousSymbolIndexes := make([][]int, 0)
	previousNumbers := make([]Numbermap, 0)

	analyzeLine := func(line string) {
		currentSymbolIndexes := regSymbols.FindAllStringIndex(line, -1)
		numberIndexes := regNumbers.FindAllStringIndex(line, -1)

		for _, numbermap := range previousNumbers {
			found := checkCurrentNumber(numbermap.index, currentSymbolIndexes)
			if found {
				sum += numbermap.num
			}
		}

		previousNumbers = nil
		symbolIndexes := append(previousSymbolIndexes, currentSymbolIndexes...)

		for _, numberIndexPair := range numberIndexes {
			number, _ := strconv.ParseInt(line[numberIndexPair[0]:numberIndexPair[1]], 10, 64)
			found := checkCurrentNumber(numberIndexPair, symbolIndexes)
			if found {
				sum += number
			} else {
				previousNumbers = append(previousNumbers, Numbermap{num: number, index: numberIndexPair})
			}
		}

		previousSymbolIndexes = currentSymbolIndexes
	}
	utils.ReadFileByLine("input.txt", analyzeLine)

	fmt.Printf("Sum: %v\n", sum)
}
