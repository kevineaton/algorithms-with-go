package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// range will ignore the nil list, so we don't need an explicit check
	for i := range list {
		if list[i] == num {
			return true
		}
	}
	return false
}
