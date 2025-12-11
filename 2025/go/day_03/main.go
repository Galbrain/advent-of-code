package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parseData(file io.Reader) (output []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return output
}

func calculatePartOne(banks []string) {
	sum := 0
	for _, bank := range banks {
		biggest := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				first := (string(bank[i]))
				second := (string(bank[j]))

				value, _ := strconv.Atoi(first + second)
				if value > biggest {
					biggest = value
				}

			}
		}
		sum += biggest
	}
	fmt.Println("Part One Sum:", sum)
}

func getBiggestDigitIndex(input string, start, end int) (index int) {
	biggestDigit := 0
	for i := start; i < end; i++ {
		elem, _ := strconv.Atoi(string(input[i]))
		if elem > biggestDigit {
			biggestDigit = elem
			index = i
		}
	}
	return index
}

func calculatePartTwo(banks []string) {
	foundTunings := []string{}
	for _, bank := range banks {
		remainder := 11
		finalTuning := ""
		windowStart := 0

		for remainder >= 0 {
			windowEnd := len(bank) - remainder

			idx := getBiggestDigitIndex(bank, windowStart, windowEnd)

			// exit early
			if idx == windowEnd-1 {
				finalTuning += bank[idx:]
				break

			}

			finalTuning += string(bank[idx])

			remainder -= 1
			windowStart = idx + 1

		}
		foundTunings = append(foundTunings, finalTuning)
	}
	sum := 0
	for _, tuning := range foundTunings {
		val, _ := strconv.Atoi(tuning)
		sum += val
	}
	fmt.Println("Part Two Sum:", sum)

}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}

	banks := parseData(file)

	calculatePartOne(banks)
	calculatePartTwo(banks)
}
