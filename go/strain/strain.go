// Package strain implement the `keep` and `discard` operation on collections.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep implement keep operation on type Ints []int
func (i Ints) Keep(f func(int) bool) Ints {
	var r Ints
	for _, e := range i {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// Discard implement keep operation on type Ints []int
func (i Ints) Discard(f func(int) bool) Ints {
	var r Ints
	for _, e := range i {
		if !f(e) {
			r = append(r, e)
		}
	}
	return r
}

// Keep implement keep operation on type Lists [][]int
func (l Lists) Keep(f func([]int) bool) Lists {
	var r Lists
	for _, e := range l {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// Keep implement keep operation on type Strings []string
func (s Strings) Keep(f func(string) bool) Strings {
	var r Strings
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}
