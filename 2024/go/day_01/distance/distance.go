package distance

func GetDistance(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

func GetSimilarity(left, right []int) (similarity int) {
	for _, a := range left {
		count := 0
		for _, b := range right {
			if a == b {
				count += 1
			}
		}
		similarity += a * count
	}

	return similarity
}
