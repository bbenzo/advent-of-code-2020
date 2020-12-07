package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_6/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	anyAnswersCounts := []int{}
	everyAnswersCounts := []int{}

	anyAnswers := make(map[byte]bool)
	peopleAnswers := make(map[byte]int)

	peopleCount := 0
	for scanner.Scan() {
		personsAnswers := scanner.Text()

		if personsAnswers == "" {
			anyAnswersCounts = append(anyAnswersCounts, len(anyAnswers))
			anyAnswers = make(map[byte]bool)

			everyCount := 0
			for _, v := range peopleAnswers {
				if v == peopleCount {
					everyCount++
				}
			}

			everyAnswersCounts = append(everyAnswersCounts, everyCount)

			peopleAnswers = make(map[byte]int)
			peopleCount = 0
			continue
		}

		peopleCount++

		for i := range personsAnswers {
			anyAnswers[personsAnswers[i]] = true
			peopleAnswers[personsAnswers[i]]++
		}
	}

	anySum := 0
	for i := range anyAnswersCounts {
		anySum += anyAnswersCounts[i]
	}

	everySum := 0
	for i := range everyAnswersCounts {
		everySum += everyAnswersCounts[i]
	}

	fmt.Printf("ADVENT OF CODE DAY 6 FIRST: %v\n", anySum)
	fmt.Printf("ADVENT OF CODE DAY 6 SECOND: %v", everySum)
}
