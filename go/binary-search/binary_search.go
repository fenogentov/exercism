package binarysearch

func SearchInts(sl []int, key int) int {
	low := 0
	high := len(sl) - 1

	for low <= high {
		median := (low + high) / 2
		if sl[median] == key {
			return median
		} else if sl[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return -1
}
