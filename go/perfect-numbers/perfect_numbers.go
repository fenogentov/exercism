package perfect

import "errors"

var ErrOnlyPositive = errors.New("only positive numbers")
var ErrUnknown = errors.New("unknown error")

type Classification string

const (
	ClassificationAbundant  Classification = "ClassificationAbundant"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationPerfect   Classification = "ClassificationPerfect"
)

func Classify(in int64) (class Classification, err error) {
	if in < 1 {
		return "", ErrOnlyPositive
	}
	var summ int64
	for i := int64(1); i < in; i++ {
		if in%i == 0 {
			summ += i
		}
	}
	switch {
	case summ == in:
		return ClassificationPerfect, nil
	case summ < in:
		return ClassificationDeficient, nil
	case summ > in:
		return ClassificationAbundant, nil
	}
	return "", ErrUnknown
}
