// Package pythagorean works with Pythagorean Triples
package pythagorean

type Triplet [3]int

// Range Finds Pythagorean Triples in a given range of value
func Range(min, max int) []Triplet {
	var triplets []Triplet

	for i := min; i <= max-2; i++ {
		for j := i + 1; j <= max-1; j++ {
			for k := i + 2; k <= max; k++ {
				if i*i+j*j == k*k {
					triplets = append(triplets, [3]int{i, j, k})
				}
			}
		}
	}
	return triplets
}

// Sum given the input integer N, finds all Pythagorean triplets for which `a + b + c = N`.
func Sum(p int) []Triplet {
	tr := Range(1, p)
	var triplets []Triplet
	for _, t := range tr {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return triplets
}
