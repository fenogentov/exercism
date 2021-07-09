// Package robotname is a solution to lerning #13 (exercism.io)
package robotname

import (
	"errors"
	"math/rand"
)

var digits = "0123456789"
var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var usedName = map[string]bool{}

// Robot is structure robot
type Robot struct {
	name string
}

// Name creates a unique name for the Robot
func (r *Robot) Name() (string, error) {
	if len(usedName) == 67700 {
		return "", errors.New("out of free names ")
	}
	buf := make([]byte, 5)
	if r.name == "" {
		for {
			buf[0] = chars[rand.Intn(len(chars))]
			buf[1] = chars[rand.Intn(len(chars))]
			buf[2] = digits[rand.Intn(len(digits))]
			buf[3] = digits[rand.Intn(len(digits))]
			buf[4] = digits[rand.Intn(len(digits))]
			name := string(buf)
			if _, ok := usedName[name]; !ok {
				usedName[name] = true
				r.name = name
				break
			}
		}
	}
	return r.name, nil
}

// Reset name of the Robot
func (r *Robot) Reset() {
	r.name = ""
}
