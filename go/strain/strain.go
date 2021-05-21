package strain

type Ints []int
type Lists [][]int
type Strings []string

func (Ints) Keep(func(int) bool) Ints {
	return nil
}

func (Ints) Discard(func(int) bool) Ints {
	return nil
}
func (Lists) Keep(func([]int) bool) Lists {
	return nil
}
func (Strings) Keep(func(string) bool) Strings {
	return nil
}
