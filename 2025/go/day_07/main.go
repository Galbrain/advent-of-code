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
	row     int
	childs  []int
	visited bool
}

func createGraph(field [][]string) (graph []Node) {
	nameCount := 0

	firstBeamLocation := slices.Index(field[0], "S")
	fmt.Println("first beam", firstBeamLocation)
	graph = append(graph, Node{name: nameCount, pos: firstBeamLocation})
	nameCount++

	fmt.Println("Starting Beam at position: ", graph)
	maxRowLen := len(field[0]) - 1

	for rIdx, row := range field {
	split:
		for sIdx, s := range row {
			if s == "^" {
				// check for nodes above splitters
				for i := len(graph) - 1; i >= 0; i-- {
					if graph[i].pos == sIdx {
						before := max(graph[i].pos-1, 0)
						after := min(graph[i].pos+1, maxRowLen)

						// add left child
						graph[i].childs = append(graph[i].childs, nameCount)
						graph = append(graph, Node{name: nameCount, pos: before, row: rIdx})
						fmt.Println(rIdx)
						nameCount++

						// add right child
						graph[i].childs = append(graph[i].childs, nameCount)
						graph = append(graph, Node{name: nameCount, pos: after, row: rIdx})
						nameCount++

						continue split
					}
				}
			}
		}
	}

	// add ending child to leafs
	for ri := range field[len(field)-1] {
		for i := len(graph) - 1; i >= 0; i-- {
			if graph[i].pos == ri && graph[i].row == len(field)-2 {
				graph[i].childs = append(graph[i].childs, -1)
			}
		}
	}
	fmt.Println("Graph: ", graph)
	sumEnd := 0
	for _, n := range graph {
		idx := slices.Index(n.childs, -1)
		if idx >= 0 {
			sumEnd++
		}
	}
	fmt.Println("sum", sumEnd)

	return graph
}

type vec2 struct {
	i int
	j int
}

func scanTree(row int, col int, field [][]string, end int, cache map[vec2]int) int {
	if row < 0 || row > end || col < 0 && col >= len(field[0]) {
		return 1
	}

	v := vec2{row, col}
	f, ok := cache[v]
	if ok {
		return f
	}

	currentPoint := field[row][col]
	result := 0
	switch currentPoint {
	case "^":
		left := scanTree(row+1, col-1, field, end, cache)
		right := scanTree(row+1, col+1, field, end, cache)
		result = left + right
	case ".":
		result = scanTree(row+1, col, field, end, cache)
	}

	cache[vec2{row, col}] = result
	return result
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

	start := slices.Index(field[0], "S")
	cache := map[vec2]int{}
	fmt.Println("len tree", len(field))
	end := len(field) - 1
	count := scanTree(1, start, field, end, cache)
	fmt.Println("Part Two - Total Timelines: ", count)
}
