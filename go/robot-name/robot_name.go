package robotname

import "math/rand"

var digits string = "0123456789"
var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Robot struct {
}

func (r Robot) Name() string {
	buf := make([]byte, 5)
	buf[0] = chars[rand.Intn(len(chars))]
	buf[1] = chars[rand.Intn(len(chars))]
	buf[2] = digits[rand.Intn(len(digits))]
	buf[3] = digits[rand.Intn(len(digits))]
	buf[4] = digits[rand.Intn(len(digits))]
	return string(buf)
}

func (r Robot) Reset() {

}
