package prime

func Factors(in int64) []int64 {
	res := []int64{}
	if in == 1 {
		return []int64{}
	}
	cnt := in
	var i int64 = 2
	for cnt > 1 {
		if cnt%i == 0 {
			res = append(res, i)
			cnt /= i
		} else {
			i++
		}
	}
	return res
}
