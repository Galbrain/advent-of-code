package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	printQueue "print_queue/internal"
	"strconv"
	"strings"
)

func parseRules(file io.Reader) (rules [][]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		split := strings.Split(str, "|")

		conv := []int{}
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error parsing rule strings", err)
			}
			conv = append(conv, i)
		}
		rules = append(rules, conv)
	}
	return rules
}

func parseUpdates(file io.Reader) (updates [][]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		split := strings.Split(str, ",")

		conv := []int{}
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error parsing update strings", err)
			}
			conv = append(conv, i)
		}
		updates = append(updates, conv)
	}
	return updates
}

func main() {
	fmt.Println("Verifying print queue.")

	rulesFile, err1 := os.Open("./data/rules.txt")
	updatesFile, err2 := os.Open("./data/updates.txt")

	if err1 != nil || err2 != nil {
		fmt.Println("Error while opening files: ", err1, err2)
		return
	}

	// parse rules
	rules := parseRules(rulesFile)

	// parse updates
	updates := parseUpdates(updatesFile)

	// check validity of updates
	checksum, fixedChecksum := printQueue.ValidateUpdates(updates, rules)
	fmt.Println("Checksum for the validity report: ", checksum)
	fmt.Println("Checksum for incorrectly ordered updates: ", fixedChecksum)

}
