package piscine

func LastWord(s string) string {

	if s == "" {
		return "\n"
	}

	word := ""
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		word = string(s[i]) + word
		i--
	}
	return word + "\n"

}
