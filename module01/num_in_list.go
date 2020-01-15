package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if list == nil || len(list) == 0 {
		return false
	}

	for i := range list {
		if list[i] == num {
			return true
		}
	}
	return false
}
