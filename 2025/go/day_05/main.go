package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(file io.Reader) (fresh []freshRange, ingreds []int) {
	scanner := bufio.NewScanner(file)

	isFresh := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isFresh = false
			continue
		}

		if isFresh {
			split := strings.Split(line, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			fresh = append(fresh, freshRange{start, end})
		} else {
			ingred, _ := strconv.Atoi(line)
			ingreds = append(ingreds, ingred)
		}
		fmt.Printf("fresh: %v, ingreds: %v\n", fresh, ingreds)
	}
	return fresh, ingreds
}

type freshRange struct {
	start int
	end   int
}

func (f freshRange) contains(num int) bool {
	if f.start <= num && num <= f.end {
		return true
	}
	return false
}

func insert(elems []freshRange, newElem freshRange, index int) []freshRange {
	return append(elems[:index], append([]freshRange{newElem}, elems[index:]...)...)
}

func consolidateFresh(fresh []freshRange) (parsedFresh []freshRange) {
fresh:
	for _, f := range fresh {
		fmt.Println("\ncurrent parsed", parsedFresh)
		fmt.Println("next to add", f)
		// first elemen2
		if len(parsedFresh) == 0 {
			parsedFresh = append(parsedFresh, f)
			continue fresh
		}

		// check existing ranges
		for i, r := range parsedFresh {
			fmt.Println("Parsed Fresh Loop", parsedFresh)
			fmt.Println("current range", r, i)

			// fully enclosed
			if r.contains(f.start) && r.contains(f.end) {
				continue fresh
			}

			// elem is in front
			if f.end < r.start {
				parsedFresh = insert(parsedFresh, f, i)
				continue fresh
			}

			// end overlaps
			if !r.contains(f.start) && r.contains(f.end) {
				parsedFresh[i].start = f.start
				continue fresh
			}

			// start overlaps
			if r.contains(f.start) && !r.contains(f.end) {
				// search for end
				j := i + 1
			next:
				for j < len(parsedFresh) {
					// end is before
					if f.end < parsedFresh[j].start {
						parsedFresh[i].end = f.end
						parsedFresh = slices.Delete(parsedFresh, i+1, j+1)
						continue fresh
					}

					// end is inside
					if parsedFresh[j].contains(f.end) {
						parsedFresh[i].end = parsedFresh[j].end
						parsedFresh = slices.Delete(parsedFresh, i+1, j+1)
						continue fresh
					}

					// end is behind, therefore continue search
					j++
					continue next
				}

				// no end found, therefore expand to last elem and delete all elems on the way
				parsedFresh[i].end = f.end
				parsedFresh = slices.Delete(parsedFresh, i+1, len(parsedFresh))
				continue fresh

			}
		}
		parsedFresh = append(parsedFresh, f)
		fmt.Println("end of step", parsedFresh)
		continue fresh
	}
	fmt.Printf("\n\nfinal consolidation: %v\n", parsedFresh)
	return parsedFresh
}

func checkTotalFreshIDs(fresh []freshRange) (sum int) {
	for _, f := range fresh {
		sum += f.end - f.start + 1
	}
	return sum
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	rawRanges, ingreds := parseInput(file)
	fmt.Printf("Raw id ranges: %v\n", rawRanges)
	fmt.Printf("Ingredients: %v\n", ingreds)

	freshRanges := consolidateFresh(rawRanges)
	fmt.Printf("Fresh id ranges: %v\n", freshRanges)

	// check ingredients
	sum := 0
	for _, item := range ingreds {
		for _, r := range freshRanges {
			if r.contains(item) {
				sum++
			}
		}
	}

	fmt.Println("Total number of fresh Ingredients: ", sum)

	totalIDs := checkTotalFreshIDs(freshRanges)
	fmt.Println("Total number of fresh IDs: ", totalIDs)
}
