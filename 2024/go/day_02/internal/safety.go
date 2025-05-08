package safety

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ReadReports(file io.Reader) (reports [][]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		report := []int{}
		for _, part := range parts {
			conv, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to int:", err)
			}

			report = append(report, conv)
		}
		reports = append(reports, report)
	}

	return reports
}

func IsReportSafe(report []int) bool {
	isIncreasing := false
	isDecreasing := false

	for i := range len(report) - 1 {
		level := report[i]
		nextLevel := report[i+1]
		diff := level - nextLevel

		if abs(diff) > 3 || abs(diff) < 1 {
			return false
		}

		if nextLevel > level {
			isIncreasing = true
			if isDecreasing {
				return false
			}
		}

		if nextLevel < level {
			isDecreasing = true
			if isIncreasing {
				return false
			}
		}

	}

	return true
}

func IsReportSafeDampened(report []int) bool {
	for i := range len(report) {
		copy := append([]int{}, report...)
		dampenedReport := slices.Delete(copy, i, i+1)

		isSafe := IsReportSafe(dampenedReport)
		if isSafe {
			return true
		}
	}

	return false
}
