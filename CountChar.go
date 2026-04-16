package piscine

func CountChar(str string, c rune) int {

	if len(str) == 0 {
		return 0
	}

	if len(str) == 0 {
		return 0
	}

	count := 0
	for _, val := range str {
		if val == c {
			count++
		}

	}
	return count
}
