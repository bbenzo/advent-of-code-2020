package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type vertex struct {
	val         string
	parentEdges []edge
	childEdges  []edge
}

type edge struct {
	parent *vertex
	child   *vertex
	weight  int
}

func main() {
	file, err := os.Open("./day_7/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	treeMap := map[string]*vertex{}
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " bags contain ")
		if len(splits) < 2 {
			log.Fatalf("invalid line in input: %vert", line)
		}

		val := splits[0]

		vert := &vertex{val: val}
		_, ok := treeMap[val]
		if strings.HasPrefix(splits[1], "contain no other") && !ok {
			treeMap[val] = vert
		}

		regex := regexp.MustCompile("\\sbag[s]*[.,]*\\s*")
		childrenStr := strings.Split(splits[1], "bag")
		for _, childStr := range childrenStr {
			weight, err := strconv.Atoi(string(childStr[0]))
			if err != nil {
				log.Fatalf("failed to parse num of child str: %vert", childStr[0])
			}

			childVal := childStr[1:]

			matchingVertex, exists := treeMap[childVal]
			if !exists {
				childVertex := &vertex{
					val: childVal,
					parentEdges: []edge{{
						parent: vert,
						child:
					}},
				}

				childVertex
			}
		}
	}
}
