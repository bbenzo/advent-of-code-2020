package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./day_5/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	maxSeatID := 0

	rows := map[int][]int{}
	minRow := math.MaxInt64
	maxRow := -1
	for _, input := range inputs {
		inputRow := row(input[:7], 0, 127)

		if inputRow < minRow {
			minRow = inputRow
		}

		if inputRow > maxRow {
			maxRow = inputRow
		}

		inputColumn := column(input[7:], 0, 7)

		rows[inputRow] = append(rows[inputRow], inputColumn)
		seatID := id(inputRow, inputColumn)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Printf("ADVENT OF CODE DAY 5 FIRST: %v\n", maxSeatID)

	for row, column := range rows {
		if len(column) < 8 && row != minRow && row != maxRow {
			sort.Ints(column)

			seat := -1
			for i := range column {
				if column[i] != i {
					seat = i
					break
				}
			}

			fmt.Printf("ADVENT OF CODE DAY 5 SECOND: %v\n", id(row, seat))
		}
	}


}

func id(inputRow int, inputColumn int) int {
	return (inputRow * 8) + inputColumn
}

func row(input string, low, high float64) int {
	if string(input[0]) == "F" {
		if len(input) == 1 {
			return int(low)
		}
		return row(input[1:], low, math.Trunc(high-((high-low)/2)))
	} else {
		if len(input) == 1 {
			return int(high)
		}
		return row(input[1:], math.RoundToEven(low+((high-low)/2)), high)
	}
}

func column(input string, low, high float64) int {
	if string(input[0]) == "L" {
		if len(input) == 1 {
			return int(low)
		}
		return column(input[1:], low, math.Trunc(high-((high-low)/2)))
	} else {
		if len(input) == 1 {
			return int(high)
		}
		return column(input[1:], math.RoundToEven(low+((high-low)/2)), high)
	}
}
