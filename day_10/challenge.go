package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day_10/sample.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	var inputs []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, val)
	}

	sort.Ints(inputs)

	fmt.Printf("ADVENT OF CODE DAY 10 FIRST: %v\n", first(inputs))
	fmt.Printf("ADVENT OF CODE DAY 10 SECOND: %v\n", second(inputs))
}

func first(inputs []int) int {
	distribution := make(map[int]int, 3)

	last := 0
	for i := 0; i < len(inputs); i++ {
		distribution[inputs[i]-last]++
		last = inputs[i]
	}

	distribution[3]++
	return distribution[1] * distribution[3]
}

func second(inputs []int) int {
	possibilities := 0

	for i := len(inputs) - 1; i > 1; i-- {
	}

	return possibilities / skipped
}
