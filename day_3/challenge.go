package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_3/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	roadmap := [][]bool{}

	count := 0
	for scanner.Scan() {
		roadmap = append(roadmap, []bool{})

		line := scanner.Text()
		for i := range line {
			if string(line[i]) == "#" {
				roadmap[count] = append(roadmap[count],true)
			} else {
				roadmap[count] = append(roadmap[count],false)
			}
		}
		count++
	}

	firstResult := traverse(roadmap, 3, 1)

	fmt.Printf("ADVENT OF CODE DAY 3 FIRST: %v\n", firstResult)

	slopes := []struct {
		right int
		down int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	secondResult := 1
	for _, slope := range slopes {
		res := traverse(roadmap, slope.right, slope.down)
		secondResult *= res
	}

	fmt.Printf("ADVENT OF CODE DAY 3 SECOND: %v", secondResult)
}

func traverse(roadmap [][]bool, right, down int) int {
	result := 0
	width := len(roadmap[0])
	rowIndex := 0

	for i := 0; i < len(roadmap) - down; i += down {
		rowIndex = (rowIndex + right) % width
		if roadmap[i+down][rowIndex] {
			result++
		}
	}
	return result
}
