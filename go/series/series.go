package series

func All(n int, s string) (out []string) {
	if n > len(s) {
		return nil
	}
	for i := n; i <= len(s); i++ {
		out = append(out, s[i-n:i])
	}

	return out
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func First(n int, s string) (out string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
