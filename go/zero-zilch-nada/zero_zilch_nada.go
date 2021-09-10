package zero

// EmptyBool returns an empty (zero value) bool
func EmptyBool() bool {
	var b bool
	return b
}

// EmptyInt returns an empty (zero value) int
func EmptyInt() int {
	var i int
	return i
}

// EmptyString returns an empty (zero value) string
func EmptyString() string {
	var s string
	return s
}

// EmptyFunc returns an empty (zero value) func
func EmptyFunc() func() {
	return nil
}

// EmptyPointer returns an empty (zero value) pointer
func EmptyPointer() *int {
	var p *int
	return p
}

// EmptyMap returns an empty (zero value) map
func EmptyMap() map[int]int {
	var m map[int]int
	return m
}

// EmptySlice returns an empty (zero value) slice
func EmptySlice() []int {
	var sl []int
	return sl
}

// EmptyChannel returns an empty (zero value) channel
func EmptyChannel() chan int {
	var ch chan int
	return ch
}

// EmptyInterface returns an empty (zero value) interface
func EmptyInterface() interface{} {
	return nil
}
