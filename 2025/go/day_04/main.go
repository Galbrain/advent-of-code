package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func parseInput(file io.Reader) (grid [][]rune) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := []rune{}
		for _, r := range line {
			parsedLine = append(parsedLine, r)
		}
		grid = append(grid, parsedLine)
	}
	return grid
}

type Pos struct {
	x int
	y int
}

func scanGrid(grid [][]rune) (validRolls int, validPositions []Pos) {
	maxHeight := len(grid)
	for i, line := range grid {
		maxWidth := len(line)

	Scan:
		for j, r := range line {
			if r != '@' {
				continue Scan
			}
			otherRolls := 0

			// scan 3x3 around current pos
		X:
			for x := i - 1; x < i+2; x++ {
				// skip out of bound
				if x < 0 || x >= maxHeight {
					continue X
				}

			Y:
				for y := j - 1; y < j+2; y++ {
					// skip out of bound
					if y < 0 || y >= maxWidth {
						continue Y
					}

					// skip if self
					if i == x && j == y {
						continue Y
					}

					if grid[x][y] == '@' {
						otherRolls++
					}
				}
			}

			if otherRolls < 4 {
				validRolls++
				pos := Pos{x: i, y: j}
				validPositions = append(validPositions, pos)
			}

		}
	}

	return validRolls, validPositions
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}

	grid := parseInput(file)
	// for _, line := range grid {
	// 	fmt.Printf("%c\n", line)
	// }

	// Part one
	foundRolls, _ := scanGrid(grid)
	fmt.Println("Part One - found valid rolls:", foundRolls)

	// Part two
	totalValidRolls := 0
	for {
		foundRolls, validPositions := scanGrid(grid)
		if foundRolls == 0 {
			break
		}

		totalValidRolls += foundRolls

		// clear valid rolls
		for _, validPos := range validPositions {
			grid[validPos.x][validPos.y] = '.'
		}
	}
	fmt.Println("Part Two - Total removed rolls:", totalValidRolls)

}
