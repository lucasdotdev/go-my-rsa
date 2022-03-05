package primes

import (
	"math"
	"math/rand"
	"time"
)

func GeneratePrime(bytes int) int {
	var min = int(math.Pow(8., float64(bytes)))

	rand.Seed(time.Now().UnixNano())

	var searchStart = rand.Intn(min)

	if searchStart%2 == 0 {
		searchStart += 1
	}

	return FindNearestPrime(searchStart + min)
}

func FindNearestPrime(max int) int {
	for possiblePrime := max; possiblePrime >= 3; possiblePrime -= 2 {
		if isPrime(possiblePrime) {
			return possiblePrime
		}
	}

	return max
}

func isPrime(target int) bool {
	if target == 0 {
		return false
	}

	if target == 1 {
		return false
	}

	if target == 3 {
		return true
	}

	if target%2 == 0 {
		return false
	}

	var root = int(math.Sqrt(float64(target))) + 3

	for possibleDivider := 3; possibleDivider <= root; possibleDivider += 2 {
		if target%possibleDivider == 0 {
			return false
		}
	}

	return true
}
