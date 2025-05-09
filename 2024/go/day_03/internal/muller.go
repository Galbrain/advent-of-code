package muller

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func getMulMatches(str string) [][]string {
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	return reg.FindAllStringSubmatch(str, -1)
}

func multiply(match []string) (int, error) {
	if len(match) != 3 {
		return 0, errors.New("Can't multiply as the input match doens't have the correct shape. Expecting array of length 3.")
	}

	num1, err1 := strconv.Atoi(match[1])
	num2, err2 := strconv.Atoi(match[2])

	if err1 != nil || err2 != nil {
		return 0, errors.New(fmt.Sprintf("Something went wrong converting string to int: \n%v, \n%v", err1, err2))
	}

	return num1 * num2, nil
}

func ExtractMulls(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	total := 0
	do := true
	for scanner.Scan() {
		rawData := scanner.Text()

		matches := getMulMatches(rawData)

		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else if do {
				res, err := multiply(match)
				if err != nil {
					fmt.Printf("Somewhting went wrong multiplying the match: %v", err)
				}
				total += res
			}
		}
	}
	return total
}
