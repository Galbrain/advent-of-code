package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseFile(file io.Reader) (total [][]string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Fields(line)
		total = append(total, splits)
	}

	return total
}

func parseFileTwo(file io.Reader) (chars [][]rune) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		currentRunes := []rune{}
		for _, c := range line {
			currentRunes = append(currentRunes, c)
		}
		chars = append(chars, currentRunes)
	}

	return chars
}

func printSlice[T any](slice []T, title string) {
	fmt.Printf("\n===== %v =====\n", title)
	for _, l := range slice {
		fmt.Println(l)
	}
	fmt.Println("")
}

func transformData(rawValues [][]string) (values [][]int, ops []string) {
	for i, rv := range rawValues {
		if i >= len(rawValues)-1 {
			// skip last row
			break
		}

		vals := []int{}
		for _, r := range rv {
			v, _ := strconv.Atoi(r)
			vals = append(vals, v)
		}
		values = append(values, vals)
	}
	ops = rawValues[len(rawValues)-1]
	return values, ops
}

func transformDataTwo(rawValues [][]rune) (words [][]string) {
	op := []string{}

	maxRow := len(rawValues) - 1
	maxCol := len(rawValues[0]) - 1

col:
	for j := maxCol; j >= 0; j-- {
		// aggregate current number
		currentWord := ""

		// walk through col
		i := 0
		for i <= maxCol {
			currRune := rawValues[i][j]

			if currRune == '+' || currRune == '*' {
				op = append(op, currentWord)
				op = append(op, string(currRune))
				words = append(words, op)
				op = []string{}
				continue col
			}

			if currRune != ' ' {
				currentWord += string(currRune)
			}

			if i == maxRow {
				op = append(op, currentWord)
				continue col
			}
			i++
		}
		op = append(op, currentWord)
	}
	return words
}

func calcTwo(rawOperation [][]string) (total int) {

opLoop:
	for _, op := range rawOperation {
		nums := []int{}

	l:
		for _, w := range op {
			if w == "+" {
				total += add(nums)
				continue opLoop
			}
			if w == "*" {
				total += multiply(nums)
				continue opLoop
			}

			num, err := strconv.Atoi(w)
			if err != nil {
				continue l
			}
			nums = append(nums, num)
		}

	}
	return total
}

func add(values []int) (sum int) {
	for _, v := range values {
		sum += v
	}
	return sum
}

func multiply(values []int) (product int) {
	product = 1
	for _, v := range values {
		product *= v
	}
	return product
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	if false {

		rawValues := parseFile(file)
		printSlice(rawValues, "Raw Values")

		values, ops := transformData(rawValues)
		// printSlice(values, "Values")
		// printSlice(ops, "ops")

		totalSum := 0
		for col := range len(values[0]) {
			currentCol := []int{}
			for row := range len(values) {
				currentCol = append(currentCol, values[row][col])
			}

			if ops[col] == "+" {
				totalSum += add(currentCol)
			} else {
				totalSum += multiply(currentCol)
			}
		}
		fmt.Println("Total sum: ", totalSum)
	}

	// === Part Two ===
	fmt.Println("\nStarting Part Two")
	rawRunes := parseFileTwo(file)

	rawOperations := transformDataTwo(rawRunes)
	totalTwo := calcTwo(rawOperations)
	fmt.Println("Total Part Two: ", totalTwo)

}
