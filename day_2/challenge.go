package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_2/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	type input struct {
		low      int
		high     int
		letter   byte
		password string
	}

	countFirst := 0
	countSecond := 0

	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		policyStrings := strings.Split(lines[0], "-")

		low, err := strconv.Atoi(policyStrings[0])
		if err != nil {
			log.Fatalf(err.Error())
		}
		high, err := strconv.Atoi(policyStrings[1])
		if err != nil {
			log.Fatalf(err.Error())
		}

		entry := input{
			low:      low,
			high:     high,
			letter:   lines[1][0],
			password: lines[2],
		}

		// first part
		letterCount := 0
		for i := range entry.password {
			if entry.password[i] == entry.letter {
				letterCount++
			}
		}

		if letterCount <= entry.high && letterCount >= low {
			countFirst++
		}

		// second part
		if (entry.password[entry.low - 1] == entry.letter || entry.password[entry.high - 1] == entry.letter) &&
			!(entry.password[entry.low - 1] == entry.letter && entry.password[entry.high - 1] == entry.letter) {
			countSecond++
		}
	}

	fmt.Printf("ADVENT OF CODE DAY 2 FIRST: %v", countFirst)
	fmt.Printf("ADVENT OF CODE DAY 2 SECOND: %v", countSecond)
}
