package module01

import (
	"fmt"
	"math"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	converted := ""
	result := dec
	iters := 0
	for result > 0 {
		nextDigit := step1(result, base)
		if nextDigit <= 9 {
			converted = fmt.Sprintf("%d%s", nextDigit, converted)
		} else {
			converted = fmt.Sprintf("%s%s", baseNumbersConversion[nextDigit], converted)
		}
		result = step2(result, base)
		iters++
		if iters > 40 {
			break
		}
	}
	return converted
}

func step1(dec, base int) int {
	result := dec % base
	return result
}

func step2(dec, base int) int {
	result := math.Floor(float64(dec) / float64(base))
	return int(result)
}

var baseNumbersConversion = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
}
