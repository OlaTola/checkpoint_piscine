package piscine

func CamelToSnakeCase(s string) string {

	if !isCamelCase(s) {
		return s
	}

	store := []rune{}
	for i, val := range s {

		//if chat is not the first item and contains capital letters, append _ before the capital char
		if i > 0 && (val >= 'A' && val <= 'Z') {
			store = append(store, '_')
		}
		//add remaing part of the str to the _
		store = append(store, val)
	}
	return string(store)

}

func isCamelCase(input string) bool {

	//if input is empty
	if input == "" {
		return false
	}

	//if first char is not alphabet
	if !(input[0] >= 'A' && input[0] <= 'Z') && !(input[0] >= 'a' && input[0] <= 'z') {
		return false
	}

	for i, val := range input {

		//if a char is capital and next xhar is also capital and index is not last char
		if (val >= 'A' && val <= 'Z') && (input[i+1] >= 'A' && input[i+1] <= 'Z') && i+1 < len(input) {
			return false
		}

		//if input contains number
		if val >= '0' && val <= '9' {
			return false
		}
	}

	//if last item is capital
	if input[len(input)-1] >= 'A' && input[len(input)-1] <= 'Z' {
		return false
	}
	return true
}
