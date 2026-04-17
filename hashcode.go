package piscine

func HashCode(dec string) string {

	if len(dec) == 0 {
		return ""
	}

	result := ""

	for _, char := range dec {
		ascii := int(char)
		hashed := (ascii + len(dec)) % 127

		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result
}
