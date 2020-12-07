package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	birthYear  int
	issueYear  int
	expireYear int
	height     string
	hairColor  string
	eyeColor   string
	passportID string
	countryID  string
}

func main() {
	file, err := os.Open("./day_4/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	valid := 0
	pass := passport{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if validatePassportID(pass.passportID) &&
				validateEyeColor(pass.eyeColor) &&
				validateHairColor(pass.hairColor) &&
				validateHeight(pass.height) &&
				pass.birthYear >= 1920 && pass.birthYear <= 2002 &&
				pass.issueYear >= 2010 && pass.issueYear <= 2020 &&
				pass.expireYear >= 2020 && pass.expireYear <= 2030 {
				valid++
			}

			pass = passport{}
			continue
		}

		splits := strings.Split(scanner.Text(), " ")

		for i := range splits {
			properties := strings.Split(splits[i], ":")

			switch properties[0] {
			case "eyr":
				intVal, err := strconv.Atoi(properties[1])
				if err != nil {
					log.Fatalf("failed to parse string val %v to integer", properties[1])
				}

				pass.expireYear = intVal
			case "byr":
				intVal, err := strconv.Atoi(properties[1])
				if err != nil {
					log.Fatalf("failed to parse string val %v to integer", properties[1])
				}

				pass.birthYear = intVal
			case "iyr":
				intVal, err := strconv.Atoi(properties[1])
				if err != nil {
					log.Fatalf("failed to parse string val %v to integer", properties[1])
				}

				pass.issueYear = intVal
			case "hgt":
				pass.height = properties[1]
			case "hcl":
				pass.hairColor = properties[1]
			case "ecl":
				pass.eyeColor = properties[1]
			case "pid":
				pass.passportID = properties[1]
			case "cid":
				pass.countryID = properties[1]
			default:
				log.Fatalf("invalid property found in input file %v", properties[0])
			}
		}
	}

	fmt.Printf("ADVENT OF CODE DAY 4 FIRST: %v", valid)
}

func validatePassportID(id string) bool {
	if id == "" || len(id) != 9 {
		return false
	}

	for _, idRune := range id {
		if idRune < 48 || idRune > 57 {
			return false
		}
	}

	return true
}

func validateEyeColor(color string) bool {
	switch color {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}

func validateHairColor(color string) bool {
	if color == "" || !strings.HasPrefix(color, "#") || len(color[1:]) != 6 {
		return false
	}

	validChars1 := "abcdef0123456789"
	for _, inputRune := range color[1:] {
		isValid := false
		for _, validRune := range validChars1 {
			if inputRune == validRune {
				isValid = true
				break
			}
		}

		if !isValid {
			return false
		}
	}

	return true
}

func validateHeight(height string) bool {
	if height == "" {
		return false
	}

	if strings.HasSuffix(height, "cm") {
		heightParts := strings.Split(height, "cm")
		heightNum, err := strconv.Atoi(heightParts[0])
		if err != nil {
			log.Fatalf("could not parse hieght str %v to integer", heightParts[0])
		}

		if heightNum >= 150 && heightNum <= 193 {
			return true
		}
	} else if strings.HasSuffix(height,"in") {
		heightParts := strings.Split(height, "in")
		heightNum, err := strconv.Atoi(heightParts[0])
		if err != nil {
			log.Fatalf("could not parse hieght str %v to integer", heightParts[0])
		}

		if heightNum >= 59 && heightNum <= 76 {
			return true
		}
	}

	return false
}
