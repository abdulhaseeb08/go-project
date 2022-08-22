package test

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, fmt.Errorf("squares out of bound")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}
