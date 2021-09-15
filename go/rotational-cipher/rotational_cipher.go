package rotationalcipher

import "strings"

func RotationalCipher(str string, shiftKey int) string {
	var res strings.Builder
	for i := range str {
		s := str[i]
		if s < 64 || (s > 90 && s < 97) || s > 122 {
			res.WriteByte(s)
			continue
		}
		b := s + byte(shiftKey)
		if s >= 64 && s <= 90 && b > 90 {
			b -= 26
		}
		if s >= 97 && s <= 122 && b > 122 {
			b -= 26
		}
		res.WriteByte(b)
	}
	return res.String()
}

rot13 := func(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}
fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))