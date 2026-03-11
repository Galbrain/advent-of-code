package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func parseData(file io.Reader) (field [][]string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, strings.Split(line, ""))
	}
	return field
}

func printField(field [][]string) {
	for _, row := range field {
		fmt.Println(row)
	}
}

type Node struct {
	name    int
	pos     int
	childs  []int
	visited bool
}

func createGraph(field [][]string) (totalSplits int) {
	graph := []Node{}

	firstBeamLocation := slices.Index(field[0], "S")
	graph = append(graph, Node{name: 0, pos: firstBeamLocation})

	fmt.Println("Starting Beam at position: ", graph)
	maxRowLen := len(field[0]) - 1

	for _, row := range field {
		// check all beam locations for splitter
		for _, node := range graph {
			if row[node.pos] == "^" {
				// calc and track beams
				totalSplits++
				before := node.pos - 1

				// safety clamp
				if before < 0 {
					before = 0
				}
				after := node.pos + 1
				if after > maxRowLen {
					after = maxRowLen
				}

				// expand graph
				node.childs = append(node.childs, before, after)
				graph = append(graph, Node{name: before})
				graph = append(graph, Node{name: after})

			}

		}
	}
	fmt.Println(graph)

	return totalSplits
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	field := parseData(file)
	printField(field)

	totalSplits := createGraph(field)
	fmt.Println("Part One - Total Splits: ", totalSplits)
}
