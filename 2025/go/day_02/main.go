package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IdRange struct {
	start int
	end   int
}

func parseInput(file io.Reader) (idRanges []IdRange, error error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for idRangeRaw := range strings.SplitSeq(line, ",") {
			elems := strings.Split(idRangeRaw, "-")
			if len(elems) != 2 {
				return nil, errors.New("Id range doesn't contain two elements")
			}

			start, err := strconv.Atoi(elems[0])
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(elems[1])
			if err != nil {
				return nil, err
			}

			idRanges = append(idRanges, IdRange{start: start, end: end})
		}
	}

	return idRanges, nil
}

func allElementsEqual(elems []string) bool {
	for i := 1; i < len(elems); i++ {
		if elems[0] != elems[i] {
			return false
		}
	}
	return true
}

func checkInvalid(idRange IdRange) (invalidIds []int) {
	fmt.Println("\n-------------------------------")
	for i := idRange.start; i <= idRange.end; i++ {
		valueString := strconv.Itoa(i)
		length := len(valueString)

		for l := 1; l <= length/2; l++ {
			if length%l != 0 {
				continue
			}

			// aggregate all possible options
			toCheck := []string{}
			for i := 0; i < length; i += l {
				toCheck = append(toCheck, valueString[i:i+l])
			}

			// validate if all items are equal
			if allElementsEqual(toCheck) && !slices.Contains(invalidIds, i) {
				invalidIds = append(invalidIds, i)
			}

		}
	}
	return invalidIds
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}

	ranges, err := parseInput(file)
	if err != nil {
		fmt.Println("Something went wrong parsing the input:", err)
		return
	}

	sum := 0
	for _, idRange := range ranges {
		invalidIds := checkInvalid(idRange)
		fmt.Println("invalidIds", invalidIds)
		for _, invalidId := range invalidIds {
			sum += invalidId
		}
	}
	fmt.Println("Sum:", sum)

}
