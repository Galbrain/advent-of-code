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

func simulateBeams(field [][]string) (totalSplits int) {

	beamLocations := make(map[int]struct{})
	firstBeamLocation := slices.Index(field[0], "S")
	beamLocations[firstBeamLocation] = struct{}{}
	fmt.Println("Starting Beam at position: ", beamLocations)
	maxRowLen := len(field[0]) - 1

	for _, row := range field {
		// check all beam locations for splitter
		for k := range beamLocations {
			if row[k] == "^" {
				totalSplits++
				before := k - 1
				if before < 0 {
					before = 0
				}
				after := k + 1
				if after > maxRowLen {
					after = maxRowLen
				}
				beamLocations[before] = struct{}{}
				beamLocations[after] = struct{}{}
				delete(beamLocations, k)
			}

		}
	}

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

	totalSplits := simulateBeams(field)
	fmt.Println("Part One - Total Splits: ", totalSplits)
}
