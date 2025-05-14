package main

import (
	"fmt"
	"os"
	xmas_search "xmas_search/internal"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
	}

	foundWords := xmas_search.SearchXmas(file)
	fmt.Println("Number of found xmas:", foundWords)
}
