package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const target = 2020

func main() {
	file, err := os.Open("./day_1/input.txt")
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs []int
	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("invalid value in input file")
		}

		inputs = append(inputs, intVal)
	}

	first, err := multiplyValuesWhichSumEqualsTarget(inputs, target)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("ADVENT OF CODE 1: %v", first)

	for i := 0; i < len(inputs); i++ {
		result, err := multiplyValuesWhichSumEqualsTarget(inputs[i+1:], target - inputs[i])
		if err == nil {
			second := inputs[i] * result
			fmt.Printf("ADVENT OF CODE 2: %v", second)
		}
	}
}

func multiplyValuesWhichSumEqualsTarget(inputs []int, target int) (int, error) {
	sort.Ints(inputs)

	left, right := 0, len(inputs) - 1

	for left < right {
		sum := inputs[left] + inputs[right]

		if sum == target {
			return inputs[left] * inputs[right], nil
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return -1, errors.New("did not find matching values")
}
