package wordy

import (
	"math"

	"regexp"

	"strconv"

	"strings"
)

func Answer(question string) (int, bool) {
	question = clean(question)

	numbers := regexp.MustCompile("-?\\d+").FindAllStringIndex(question, -1)
	operations := regexp.MustCompile("[a-z]+( [a-z]+)*").FindAllStringIndex(question, -1)

	if len(operations)+1 != len(numbers) {
		return 0, false
	}

	if len(operations) > 0 && operations[0][0] < numbers[0][0] {
		return 0, false
	}

	acc, _ := strconv.Atoi(substring(question, numbers[0]))

	for i, o := range operations {
		if o[0] > numbers[i+1][0] {
			return 0, false
		}

		operation := substring(question, o)
		number, _ := strconv.Atoi(substring(question, numbers[i+1]))

		switch operation {
		case "plus":
			acc += number
		case "minus":
			acc -= number
		case "multiplied by":
			acc *= number
		case "divided by":
			acc /= number
		case "raised to the":
			acc = pow(acc, number)
		default:
			return 0, false
		}
	}

	return acc, true
}

func substring(s string, indexes []int) string {
	return s[indexes[0]:indexes[1]]
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func clean(question string) string {
	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, "?", "")
	return regexp.MustCompile("[a-z]{2} power").ReplaceAllString(question, "")
}
