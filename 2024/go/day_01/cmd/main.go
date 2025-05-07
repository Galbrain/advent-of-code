package main

import (
	"distance/distance"
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	filePath := flag.String("file", "input_short.txt", "Input file path")
	flag.Parse()

	// read file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}

	left, right, err := distance.ExtractPairs(file)

	// sort pairs
	slices.Sort(left)
	slices.Sort(right)

	// calc distance
	totalDist := 0
	for i, a := range left {
		totalDist += distance.GetDistance(a, right[i])
	}
	fmt.Println("Total distance:", totalDist)

	// calc similarity
	fmt.Println(len(left), len(right))
	similarity := distance.GetSimilarity(left, right)
	fmt.Println("Total similarity:", similarity)

}
