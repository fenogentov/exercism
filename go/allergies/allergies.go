package allergies

var allergens map[uint]string = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) []string {
	var i uint
	var res []string
	for i = 1; i <= score; i = i << 1 {
		n := score & i
		if a, ok := allergens[n]; ok {
			res = append(res, a)
		}
	}
	return res
}

func AllergicTo(score uint, substance string) bool {
	algn := Allergies(score)
	for _, a := range algn {
		if a == substance {
			return true
		}
	}
	return false
}
