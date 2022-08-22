package test

import (
	"strings"
)

func ToRNA(dna string) string {

	if dna == "" {
		return ""
	}
	splitDNAstring := strings.Split(dna, "")
	newString := ""
	for _, v := range splitDNAstring {
		switch {
		case v == "G":
			newString += "C"
		case v == "C":
			newString += "G"
		case v == "T":
			newString += "A"
		case v == "A":
			newString += "U"
		}
	}
	return newString
}
