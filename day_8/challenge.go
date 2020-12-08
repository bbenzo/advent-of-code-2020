package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	noOperation string = "nop"
	jump               = "jmp"
	accumulate         = "acc"
)

type instruction struct {
	name string
	val  int
}

func main() {
	file, err := os.Open("day_8/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	var instructions []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to parse line")
		}

		instructions = append(instructions, instruction{
			name: parts[0],
			val:  val,
		})
	}

	accumulator := run(instructions)

	fmt.Printf("Program exited: %v", accumulator)
}

func run(instructions []instruction) int {
	visited := map[int]bool{}
	accumulator := 0
	i := 0
	lastSwitchIndex := -1
	switchIndex := 0
	for i < len(instructions) {
		if visited[i] {
			// switch last one back
			if lastSwitchIndex >= 0 {
				switchInstruction(&instructions[lastSwitchIndex])
			}

			// switch next one
			switchInstruction(&instructions[switchIndex])
			switchIndex++

			// save switched one
			lastSwitchIndex = switchIndex - 1

			// reset visited and start
			visited = map[int]bool{}
			i = 0
			accumulator = 0
		}

		visited[i] = true
		instru := instructions[i]

		switch instru.name {
		case accumulate:
			accumulator += instru.val
		case jump:
			i = i + instru.val
			continue
		}
		i++
	}
	return accumulator
}

func switchInstruction(instruction *instruction) bool {
	switch instruction.name {
	case noOperation:
		instruction.name = jump
		return true
	case jump:
		instruction.name = noOperation
		return true
	}

	return false
}
