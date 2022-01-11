package secret

func Handshake(code uint) []string {

	sec := []string{}
	if (code & 8) == 8 {
		sec = append(sec, "jump")
	}
	if (code & 4) == 4 {
		sec = append(sec, "close your eyes")
	}
	if (code & 2) == 2 {
		sec = append(sec, "double blink")
	}
	if (code & 1) == 1 {
		sec = append(sec, "wink")
	}

	if (code & 16) == 16 {
		return sec
	}

	rev := make([]string, len(sec))
	for i, s := range sec {
		rev[len(sec)-i-1] = s
	}
	return rev
}
