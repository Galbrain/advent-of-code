package main

import (
	"bufio"
	"fmt"
	"guard/guard"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Start Guard path analysis.")

	// read data
	file, err := os.Open("./data/input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	// parse data
	area := parseMap(file)

	totalPositions := guard.PredictGuardPath(area)
	fmt.Println("Amount of spots visited by guard:", totalPositions)

}

func parseMap(file io.Reader) (area [][]string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		c := strings.Split(str, "")
		area = append(area, c)
	}
	guard.PrintArea(area)

	return area
}
