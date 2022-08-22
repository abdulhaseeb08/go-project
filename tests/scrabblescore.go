package test

import (
	"strings"
)

func Score(word string) int {
	var count int

	word = strings.ToUpper(word)

	sliceOfString := strings.Split(word, "")

	for _, v := range sliceOfString {
		switch {
		case v == "A" || v == "E" || v == "I" || v == "O" || v == "U" || v == "L" || v == "N" || v == "R" || v == "S" || v == "T":
			count += 1
		case v == "D" || v == "G":
			count += 2
		case v == "B" || v == "C" || v == "M" || v == "P":
			count += 3
		case v == "F" || v == "H" || v == "V" || v == "W" || v == "Y":
			count += 4
		case v == "K":
			count += 5
		case v == "J" || v == "X":
			count += 8
		case v == "Q" || v == "Z":
			count += 10
		}
	}
	return count
}
