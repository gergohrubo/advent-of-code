package main

import (
	"aoc/2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

func checkPreviousGears(previousGears []GearMap, numberIndexes [][]int, line string) int64 {
	var sum int64 = 0
	for _, gearMap := range previousGears {
		for _, numberIndexPair := range numberIndexes {
			if gearMap.index >= numberIndexPair[0]-1 && gearMap.index <= numberIndexPair[1] {
				number, _ := strconv.ParseInt(line[numberIndexPair[0]:numberIndexPair[1]], 10, 64)
				gearMap.numbers = append(gearMap.numbers, number)
			}
		}
		if len(gearMap.numbers) == 2 {
			sum += gearMap.numbers[0] * gearMap.numbers[1]
		}
	}
	return sum
}

func partTwo() {
	var sum int64 = 0
	regNumbers := regexp.MustCompile(`([1-9][0-9]*)`)
	regGear := regexp.MustCompile(`\*`)

	previousGears := make([]GearMap, 0)
	previousNumbers := make([]Numbermap, 0)

	analyzeLine := func(line string) {
		numberIndexes := regNumbers.FindAllStringIndex(line, -1)
		gearIndexes := regGear.FindAllStringIndex(line, -1)

		partSum := checkPreviousGears(previousGears, numberIndexes, line)
		sum += partSum

		previousGears = nil

		for _, gearIndexPair := range gearIndexes {
			gearNumbers := make([]int64, 0)

			for _, numberMap := range previousNumbers {
				if gearIndexPair[0] >= numberMap.index[0]-1 && gearIndexPair[0] <= numberMap.index[1] {
					gearNumbers = append(gearNumbers, numberMap.num)
				}
			}

			for _, numberIndexPair := range numberIndexes {
				if gearIndexPair[0] == numberIndexPair[0]-1 || gearIndexPair[0] == numberIndexPair[1] {
					number, _ := strconv.ParseInt(line[numberIndexPair[0]:numberIndexPair[1]], 10, 64)
					gearNumbers = append(gearNumbers, number)
				}
			}

			previousGears = append(previousGears, GearMap{index: gearIndexPair[0], numbers: gearNumbers})
		}

		previousNumbers = nil

		for _, numberIndexPair := range numberIndexes {
			number, _ := strconv.ParseInt(line[numberIndexPair[0]:numberIndexPair[1]], 10, 64)
			previousNumbers = append(previousNumbers, Numbermap{num: number, index: numberIndexPair})
		}
	}
	utils.ReadFileByLine("input.txt", analyzeLine)

	fmt.Printf("Part two Sum: %v\n", sum)
}
