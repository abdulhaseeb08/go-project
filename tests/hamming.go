package test

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {

	sliceOfa := strings.Split(a, "")
	sliceOfb := strings.Split(b, "")

	if len(sliceOfa) != len(sliceOfb) {
		return 0, errors.New("unequal lengths")
	}

	var count int

	for i, v := range sliceOfa {
		if sliceOfb[i] != v {
			count++
		}
	}
	return count, nil

}
