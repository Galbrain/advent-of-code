package distance

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Pair struct {
	A, B int
}

func ExtractPairs(file io.Reader) (left, right []int, err error) {
	fmt.Println("extracting pairs of file:", file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format, expected two numbers per line:", line)
			return left, right, err
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to int:", err1, err2)
			return left, right, err
		}

		left = append(left, num1)
		right = append(right, num2)

	}

	return left, right, nil
}
