package piscine

func FirstWord(s string) string {
	firstword := ""
	for _, val := range s {
		if val == ' ' {
			break
		}
		firstword += string(val)
	}
	return firstword + "\n"
}
