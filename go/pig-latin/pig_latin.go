package piglatin

import (
	"strings"
)

var vowelSound = map[byte]interface{}{'a': nil, 'u': nil, 'e': nil, 'i': nil, 'o': nil}

func Sentence(str string) string {
	sl := strings.Split(str, " ")
	for i, s := range sl {
		sl[i] = encryption(s)
	}
	return strings.Join(sl, " ")
}

func encryption(str string) string {
	if _, ok := vowelSound[str[0]]; ok {
		str += "ay"
		return str
	}
	if str[:2] == "yt" || str[:2] == "xr" {
		str += "ay"
		return str
	}
	if _, ok := vowelSound[str[0]]; !ok {
		for {
			if _, ok := vowelSound[str[0]]; !ok {
				if str[:2] == "qu" {
					str = str[1:] + string(str[0])
				}
				str = str[1:] + string(str[0])
				if str[0] == 'y' {
					break
				}
			} else {
				break
			}
		}
		str += "ay"
		return str
	}
	return ""
}
