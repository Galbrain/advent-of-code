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

type cube struct {
	color string
	count int
}

type round struct {
	cubes []cube
}

type game struct {
	id     int
	rounds []round
}

func processGamesData(file io.Reader) (games []game) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		game := game{}

		gameAndRounds := strings.Split(line, ":")
		fields := strings.Fields(gameAndRounds[0])

		// extract id
		id, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error converting game id to int: ", err)
			return
		}
		game.id = id

		// extract rounds
		rawRounds := strings.SplitSeq(gameAndRounds[1], ";")
		for rawRound := range rawRounds {
			round := round{}

			rawCubes := strings.SplitSeq(rawRound, ",")
			for rawCube := range rawCubes {
				split := strings.Fields(rawCube)
				count, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println("Error converting cube count: ", err)
				}
				cube := cube{count: count, color: split[1]}
				round.cubes = append(round.cubes, cube)
			}
			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}

	return games
}

func validateGame(game game) bool {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	totalBag := 0
	for _, c := range bag {
		totalBag += c
	}

	for _, round := range game.rounds {
		total := 0
		for _, cube := range round.cubes {
			total += cube.count
			if cube.count > bag[cube.color] {
				return false
			}
		}

		if total > totalBag {
			return false
		}
	}

	return true
}

func findMinimumSet(game game) (power int) {
	minimumSet := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range game.rounds {
		for _, cube := range round.cubes {
			if cube.count > minimumSet[cube.color] {
				minimumSet[cube.color] = cube.count
			}
		}
	}

	power = 1
	for _, s := range minimumSet {
		power *= s
	}

	return power
}

func main() {
	filePath := flag.String("file", "input.txt", "Input file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	games := processGamesData(file)
	if false {
		if len(games) <= 5 {
			fmt.Printf("Found and processed games:\n%v\n", games)
		} else {
			fmt.Println("Found and processed games, only showing length: ", len(games))
		}
	}

	sum := 0
	power := 0
	for _, game := range games {
		if validateGame(game) {
			sum += game.id
		}

		power += findMinimumSet(game)

	}
	fmt.Println("Total sum:", sum)
	fmt.Println("Total power:", power)

}
