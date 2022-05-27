package say

func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}

	s := convert(n)

	return s, true
}

func convert(n int64) string {
	if n == 0 {
		return "zero"
	} else if n <= 19 {
		return []string{"one", "two", "three", "four", "five", "six", "seven", "eight",
			"nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
			"seventeen", "eighteen", "nineteen"}[n-1]
	} else if n <= 99 {
		s := ""
		if n%10 != 0 {
			s = "-" + convert(n%10)
		}
		return []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy",
			"eighty", "ninety"}[n/10-2] + s
	} else if n <= 199 {
		s := ""
		if n%100 != 0 {
			s = " " + convert(n%100)
		}
		return "one hundred" + s
	} else if n <= 999 {
		s := ""
		if n%100 != 0 {
			s = " " + convert(n%100)
		}
		return convert(n/100) + " hundred" + s
	} else if n <= 1999 {
		s := ""
		if n%1000 != 0 {
			s = " " + convert(n%1000)
		}
		return "one thousand" + s
	} else if n <= 999999 {
		s := ""
		if n%1000 != 0 {
			s = " " + convert(n%1000)
		}
		return convert(n/1000) + " thousand" + s
	} else if n <= 1999999 {
		s := ""
		if n%1000000 != 0 {
			s = " " + convert(n%1000000)
		}
		return "one million" + s
	} else if n <= 999999999 {
		s := ""
		if n%1000000 != 0 {
			s = " " + convert(n%1000000)
		}
		return convert(n/1000000) + " million" + s
	} else if n <= 1999999999 {
		s := ""
		if n%1000000000 != 0 {
			s = " " + convert(n%1000000000)
		}
		return "one billion" + s
	} else {
		s := ""
		if n%1000000000 != 0 {
			s = " " + convert(n%1000000000)
		}
		return convert(n/1000000000) + " billion" + s
	}
}
