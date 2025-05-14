package xmas_search

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func convertToSlice(file io.Reader) (matrix [][]string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		split := strings.Split(str, "")

		matrix = append(matrix, split)
	}
	return matrix
}

func isSymmetric(input [][]string) bool {
	width := len(input[0])
	for i, line := range input {
		if i == 0 {
			continue
		}

		if len(line) != width {
			return false
		}
	}

	return true
}

func extractInDirection(input [][]string, search string, i, j, di, dj int) bool {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
		return false
	}

	if di < -1 || di > 1 || dj < -1 || dj > 1 {
		fmt.Println("The direction inputs must be between including -1 and 1.")
		return false
	}

	boundI := i + (len(search)-1)*di
	boundJ := j + (len(search)-1)*dj

	if boundI < 0 || boundJ < 0 || boundI >= len(input) || boundJ >= len(input[i]) {
		return false
	}

	for s := range len(search) {
		if string(search[s]) != input[i][j] {
			return false
		}

		i += di
		j += dj
	}

	return true

}

func extractWord(input [][]string) (count int, err error) {
	if len(input) == 0 {
		return count, err
	}

	if !isSymmetric(input) {
		return count, fmt.Errorf("Input is not symmetric, therefore stopping.")
	}

	searchTerm := "MAS"

	for i, line := range input {
		if len(input) == 0 {
			break
		}
		for j := range line {
			if input[i][j] != "A" {
				continue
			}

			diagFound := 0

			// north east
			if extractInDirection(input, searchTerm, i+1, j-1, -1, 1) {
				diagFound += 1
			}

			if extractInDirection(input, searchTerm, i-1, j-1, 1, 1) {
				diagFound += 1
			}

			if extractInDirection(input, searchTerm, i-1, j+1, 1, -1) {
				diagFound += 1
			}

			if extractInDirection(input, searchTerm, i+1, j+1, -1, -1) {
				diagFound += 1
			}

			if diagFound == 2 {
				count += 1
			}

		}
	}

	return count, nil
}

func SearchXmas(file io.Reader) int {
	matrix := convertToSlice(file)

	count, err := extractWord(matrix)
	if err != nil {
		fmt.Println("Error while extracting xmas:", err)
		return count
	}

	return count
}
