package railfence

import "strings"

func Encode(str string, n int) string {
	if len(str) < 2 {
		return str
	}
	var rez strings.Builder
	// select first row
	for i := 0; i < len(str); i += (n - 1) * 2 {
		rez.WriteByte(str[i])
	}
	// select middle rows
	for j := 1; j < n-1; j++ {
		down := true
		for i := j; i < len(str); {
			rez.WriteByte(str[i])
			if down {
				i += (n - j - 1) * 2
			} else {
				i += (n-1)*2 - (n-j-1)*2
			}

			down = !down
		}
	}
	// select last row
	for i := n - 1; i < len(str); i += (n - 1) * 2 {
		rez.WriteByte(str[i])
	}
	return rez.String()
}

func Decode(str string, n int) string {

	// detect row lengths
	l := make([]int, n)
	down := false
	var row int
	for i := 0; i < len(str); i++ {
		l[row]++
		if down {
			row--
		} else {
			row++
		}
		if row == 0 || row == n-1 {
			down = !down
		}
	}
	// fill the rows
	rows := make([]string, n)
	inp := str
	for i := 0; i < n; i++ {
		rows[i] = inp[:l[i]]
		inp = inp[l[i]:]
	}
	// zig-zag
	var ret strings.Builder
	down = false
	row = 0
	for i := 0; i < len(str); i++ {
		ret.WriteByte(rows[row][0])
		rows[row] = rows[row][1:]
		if down {
			row--
		} else {
			row++
		}
		if row == 0 || row == n-1 {
			down = !down
		}
	}
	return ret.String()
}
