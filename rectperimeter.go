package piscine

func RectPerimeter(w, h int) int {

	per := 0
	if w > 0 && h > 0 {
		per = 2*w + 2*h
		return per
	} else {
		return -1
	}

}
