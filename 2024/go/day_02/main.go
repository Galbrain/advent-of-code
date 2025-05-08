package main

import (
	"fmt"
	"os"
	safety "safety/internal"
)

func main() {
	fmt.Println("Initializing safety levels analysis.")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}

	reports := safety.ReadReports(file)

	totalSafe := 0
	for _, report := range reports {
		isSafe := safety.IsReportSafeDampened(report)
		if isSafe {
			totalSafe += 1
		}
	}

	fmt.Printf("The total number of safe reports is: %d\n", totalSafe)
}
