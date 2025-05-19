package guard

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func input(label string) {
	var s string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+"\n")
		s, _ = reader.ReadString('\n')
		if s != "" {
			break
		}
	}
}

func PredictGuardPath(area [][]string) (totalPositions int) {
	visitedSpots := []string{}

	posX, posY := getCurrentGuardPosition(area)
	direction := North
	fmt.Println("Guard starting at:", posX, posY)

	guardInArea := true
	for guardInArea {
		// input("Press ENTER to continue..")
		posString := strconv.Itoa(posX) + "," + strconv.Itoa(posY)
		if slices.Index(visitedSpots, posString) == -1 {
			visitedSpots = append(visitedSpots, posString)
		}

		nextX, nextY := move(posX, posY, direction)
		if checkIfGuardaLeft(area, nextX, nextY) {
			area[posY][posX] = "X"
			break
		}

		newDir, ok := checkObstacle(area, nextX, nextY, direction)
		if !ok {
			nextX, nextY = move(posX, posY, newDir)
		}

		area[posY][posX] = "X"
		area[nextY][nextX] = "^"

		posX, posY = nextX, nextY
		direction = newDir
		if checkIfGuardaLeft(area, nextX, nextX) {
			break
		}

		area[posY][posX] = getGuardSymbol(newDir)
	}

	totalPositions = len(visitedSpots)
	return totalPositions
}

func getCurrentGuardPosition(area [][]string) (posX, posY int) {
	for y, row := range area {
		x := slices.Index(row, "^")
		if x > -1 {
			return x, y
		}
	}

	return posX, posY
}

func checkIfGuardaLeft(area [][]string, x, y int) bool {
	if x < 0 || y < 0 || x >= len(area[0]) || y >= len(area) {
		return true
	}
	return false
}

func getGuardSymbol(dir Direction) string {
	switch dir {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	default:
		panic("No valid direction was given")

	}
}

func move(x, y int, direction Direction) (posX, posY int) {
	switch direction {
	case North:
		y -= 1
	case East:
		x += 1
	case South:
		y += 1
	case West:
		x -= 1
	default:
		panic("AHHH")
	}

	return x, y
}

// checks if next field is a '#', if so will return a new direction to move in
// Also returns "ok" as true if no obstacle was found.
func checkObstacle(area [][]string, x, y int, dir Direction) (direction Direction, ok bool) {
	if area[y][x] == "#" {
		switch dir {
		case North:
			return East, false
		case East:
			return South, false
		case South:
			return West, false
		case West:
			return North, false
		default:
			panic("No valid direction was input for checkObstacle.")
		}
	}

	return dir, true
}

type LoopCandidate struct {
	x        int
	y        int
	poitions []string
}

func searchLoops(area [][]string) int {

	return -1
}

func PrintArea(area [][]string) {
	fmt.Println("----------------------------")
	for _, row := range area {
		fmt.Println(strings.Join(row, ""))
	}
}
