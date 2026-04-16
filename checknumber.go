package piscine

func CheckNumber(args string) bool {
	for _, val := range args {
		if val >= '0' && val <= '9' {
			return true
		}
	}
	return false
}
