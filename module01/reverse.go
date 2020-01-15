package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// note that for large strings, there are more efficient string build options
func Reverse(word string) string {
	reversed := ""
	for _, r := range word {
		reversed = string(r) + reversed
	}
	return reversed
}
