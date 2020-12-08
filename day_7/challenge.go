package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const shinyGold = "shiny gold"

type graph struct {
	vertices map[string]*vertex
}

type vertex struct {
	val         string
	parentEdges []*edge
	childEdges  []*edge
}

type edge struct {
	from   *vertex
	to     *vertex
	weight int
}

func main() {
	file, err := os.Open("./day_7/input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)

	graph := graph{vertices: map[string]*vertex{}}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		splits := strings.Split(line, " bags contain ")
		if len(splits) < 2 {
			log.Fatalf("invalid line in input: %vert", line)
		}

		val := splits[0]

		vert := setVertex(graph, val)

		if strings.HasPrefix(splits[1], "no other") {
			continue
		}

		for _, childStr := range childStrings(splits) {
			if childStr == "" {
				continue
			}

			weight := parseWeight(childStr)

			childVal := parseVertexName(childStr)

			match, exists := graph.vertices[childVal]
			if !exists {
				// create vertex
				match = &vertex{
					val: childVal,
				}

				graph.vertices[match.val] = match
			}

			// create parent edge
			parentEdge := &edge{
				from:   match,
				to:     vert,
				weight: weight,
			}

			// create child edge
			childEdge := &edge{
				from:   vert,
				to:     match,
				weight: weight,
			}

			match.parentEdges = append(match.parentEdges, parentEdge)
			vert.childEdges = append(vert.childEdges, childEdge)
		}

	}

	calculateParentEdges(graph.vertices[shinyGold].parentEdges)
	fmt.Printf("ADVENT OF CODE DAY 7 FIRST: %v\n", len(hits))

	result := calculateChildEdges(graph.vertices[shinyGold], 0)
	fmt.Printf("ADVENT OF CODE DAY 7 SECOND: %v\n", result)
}

var hits = make(map[string]bool)

func calculateParentEdges(edges []*edge) {
	for _, edge := range edges {
		hits[edge.to.val] = true
		calculateParentEdges(edge.to.parentEdges)
	}
}

func calculateChildEdges(v *vertex, count int) int {
	for _, edge := range v.childEdges {
		count++
		count *= edge.weight
		count = calculateChildEdges(edge.to, count)
	}

	return count
}

func parseVertexName(childStr string) string {
	childVal := childStr[2:]
	return childVal
}

func parseWeight(childStr string) int {
	weight, err := strconv.Atoi(string(childStr[0]))
	if err != nil {
		log.Fatalf("failed to parse num of child str: %vert", childStr[0])
	}
	return weight
}

func childStrings(splits []string) []string {
	regex := regexp.MustCompile("\\sbag[s]*[.,]*\\s*")

	childrenStr := regex.Split(splits[1], -1)
	return childrenStr
}

func setVertex(graph graph, val string) *vertex {
	vert, ok := graph.vertices[val]
	if !ok {
		vert = &vertex{val: val}
		graph.vertices[val] = vert
	}
	return vert
}
