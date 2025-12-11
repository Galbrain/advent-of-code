package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parseInputData(file io.Reader) (operations []string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		operations = append(operations, line)
	}
	return operations
}

func rotateDial(position int, operation string) (finalPosition int, passedZero int) {
	direction := 0
	if operation[0] == 'R' {
		direction = 1
	}

	operation = operation[1:]
	steps, err := strconv.Atoi(operation)
	if err != nil {
		println("Error while converting operation to steps: ", err)
	}

	for range steps {
		if direction == 0 {
			position -= 1
			if position == -1 {
				position = 99
			}
		} else {
			position += 1
			if position == 100 {
				position = 0
			}
		}

		if position == 0 {
			passedZero += 1
		}
	}

	return position, passedZero
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}

	operations := parseInputData(file)

	position := 50
	totalStoppedAtZero := 0
	for _, op := range operations {
		passedZero := 0
		position, passedZero = rotateDial(position, op)
		totalStoppedAtZero += passedZero
	}

	println("Dial stopped at zero: ", totalStoppedAtZero)
}
