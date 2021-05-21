// Package twofer is a solution to lerning #2 (exercism.io)
package twofer

// ShareWith return message, given a name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
