package test

func Sieve(limit int) []int {
	var primeList = []int{0, 0}

	for i := 2; i <= limit; i++ {
		primeList = append(primeList, i)
	}

	for i := 0; i < limit-1; i++ {
		if primeList[i] != 0 {
			for j := i * i; j <= limit; j += i {
				primeList[i] = 0
			}
		}
	}

	return primeList
}
