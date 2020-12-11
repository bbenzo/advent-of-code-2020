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
	file, err := os.Open("day_9/input.txt")
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

	firstResult, index := first(25, inputs)
	fmt.Printf("ADVENT OF CODE DAY 9 FIRST: %v\n", firstResult)

	result := second(firstResult, inputs[:index])
	fmt.Printf("ADVENT OF CODE DAY 9 SECOND: %v\n", result)
}

func second(fromFirst int, arr []int) int {

	sum := arr[0]
	start, end := 0, 1
	for start < end {
		for sum > fromFirst {
			sum -= arr[start]
			start++
		}

		if sum == fromFirst {
			arr = arr[start:end]
			sort.Ints(arr)

			return arr[0] + arr[len(arr)-1]
		}

		sum += arr[end]

		end++
	}
	return -1
}

func first(preambleLen int, arr []int) (int, int) {
	for i := preambleLen; i < len(arr); i++ {
		observed := arr[i]

		tmp := make([]int, preambleLen)
		copy(tmp, arr[i-preambleLen:i])
		sort.Ints(tmp)

		found := false
		start, end := 0, preambleLen-1
		for start < end {
			sum := tmp[start] + tmp[end]

			if sum == observed {
				found = true
				break
			}

			if sum < observed {
				start++
			} else {
				end--
			}
		}

		if !found {
			return observed, i
		}
	}

	return -1, -1
}
