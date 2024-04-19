package main

import (
	"aoc/2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func partOne() {
	var sum int64 = 0

	limitMap := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	analyzeLine := func(line string) {
		res := strings.Split(line, ":")
		if len(res) > 2 {
			fmt.Print(line)
			log.Fatal("Incorrect line format")
		}
		gameId, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(strings.TrimSpace(res[0]), "Game")[1]), 10, 64)
		draws := strings.Split(res[1], ";")
		for _, draw := range draws {
			items := strings.Split(draw, ",")
			for _, item := range items {
				itemSplit := strings.Split(strings.TrimSpace(item), " ")
				color := itemSplit[1]
				amount, _ := strconv.ParseInt(itemSplit[0], 10, 64)
				if limitMap[color] < amount {
					return
				}
			}
		}
		sum += gameId
	}
	utils.ReadFileByLine("input.txt", analyzeLine)
	fmt.Printf("Sum: %v\n", sum)
}

func partTwo() {
	var sum int64 = 0

	analyzeLine := func(line string) {
		fmt.Printf("Line: %v\n", line)
		minMap := map[string]int64{
			"red":   -1,
			"green": -1,
			"blue":  -1,
		}
		var power int64 = 1

		res := strings.Split(line, ":")
		if len(res) > 2 {
			fmt.Print(line)
			log.Fatal("Incorrect line format")
		}
		draws := strings.Split(res[1], ";")
		for _, draw := range draws {
			items := strings.Split(draw, ",")
			for _, item := range items {
				itemSplit := strings.Split(strings.TrimSpace(item), " ")
				color := itemSplit[1]
				amount, _ := strconv.ParseInt(itemSplit[0], 10, 64)
				if minMap[color] == -1 || minMap[color] < amount {
					minMap[color] = amount
				}
			}
		}
		for color, amount := range minMap {
			fmt.Printf("Color: %v\n", color)
			fmt.Printf("Amount: %v\n", strconv.Itoa(int(amount)))
			power *= amount
		}
		sum += power
	}
	utils.ReadFileByLine("input.txt", analyzeLine)
	fmt.Printf("Sum: %v\n", sum)
}

func main() {
	// partOne()
	partTwo()
}
