package main

import (
	"fmt"
	muller "mull/internal"
	"os"
)

func main() {
	fmt.Println("Initializing analysis of corrupted data..")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
	}

	// parse data
	total := muller.ExtractMulls(file)
	fmt.Println("total:", total)

}
