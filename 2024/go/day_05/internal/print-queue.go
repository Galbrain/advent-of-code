package printQueue

import (
	"slices"
)

func ValidateUpdates(updates [][]int, rules [][]int) (checksum int, fixedChecksum int) {
	invalidUpdates := [][]int{}
	for _, update := range updates {
		isValid := validateUpdate(update, rules)
		if !isValid {
			invalidUpdates = append(invalidUpdates, update)
			continue
		}

		middle := (len(update) - 1) / 2
		checksum += update[middle]
	}

	for _, update := range invalidUpdates {
		fixedUpdate := fixBrokenUpdate(update, rules)
		middle := (len(fixedUpdate) - 1) / 2
		fixedChecksum += fixedUpdate[middle]
	}

	return checksum, fixedChecksum
}

func getMatchingRules(rules [][]int, page int) (matchingRules [][]int) {
	for _, rule := range rules {
		if slices.Contains(rule, page) {
			matchingRules = append(matchingRules, rule)
		}
	}
	return matchingRules
}

func validateUpdate(update []int, rules [][]int) bool {
	for pi, page := range update {

		// get matching rules
		matchingRules := getMatchingRules(rules, page)

		for _, rule := range matchingRules {
			// check first position rules
			if rule[0] == page {
				i := slices.Index(update, rule[1])
				if i >= 0 && i < pi {
					return false
				}
			}

			// check second position rules
			if rule[1] == page {
				i := slices.Index(update, rule[0])
				if i >= 0 && i > pi {
					return false
				}
			}

		}

	}
	return true
}

func fixBrokenUpdate(update []int, rules [][]int) []int {
	fullyFixed := false

	for !fullyFixed {
		for pi, page := range update {

			// get matching rules
			matchingRules := getMatchingRules(rules, page)

			for _, rule := range matchingRules {
				// check first position rules
				if rule[0] == page {
					i := slices.Index(update, rule[1])
					if i >= 0 && i < pi {
						update[pi], update[i] = update[i], update[pi]
						break
					}
				}

				// check second position rules
				if rule[1] == page {
					i := slices.Index(update, rule[0])
					if i >= 0 && i > pi {
						update[pi], update[i] = update[i], update[pi]
						break
					}
				}

			}

		}

		if validateUpdate(update, rules) {
			fullyFixed = true
		}

	}

	return update
}
