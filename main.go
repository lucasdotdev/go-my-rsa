package main

import (
	"math/big"
	"my-rsa/primes"
)

type profile struct {
	p          int
	q          int
	n          int
	t          int
	e          int
	d          int
	publicKey  key
	privateKey key
}

type key struct {
	firstPiece  int
	secondPiece int
}

func main() {
}

func LoadProfile(p, q, n, t, e, d int) profile {
	var publicKey = key{n, e}
	var privateKey = key{n, d}

	return profile{
		p,
		q,
		n,
		t,
		e,
		d,
		publicKey,
		privateKey,
	}
}

func GenerateProfile(bytes int) profile {
	var p = primes.GeneratePrime(bytes)
	var q = primes.GeneratePrime(bytes)

	for q == p {
		q = primes.GeneratePrime(bytes)
	}

	var n = calculateN(p, q)
	var t = calculateT(p, q)
	var e = primes.GeneratePrime(bytes)
	var d = generateD(e, t)

	var publicKey = key{n, e}
	var privateKey = key{n, d}

	return profile{
		p,
		q,
		n,
		t,
		e,
		d,
		publicKey,
		privateKey,
	}
}

func EncryptInt(target int, base profile) int {
	return processIntWithKey(target, base.publicKey)
}

func DecryptInt(target int, base profile) int {
	return processIntWithKey(target, base.privateKey)
}

func EncryptString(target string, base profile) string {
	return processStringWithKey(target, base.publicKey)
}

func DecryptString(target string, base profile) string {
	return processStringWithKey(target, base.privateKey)
}

func processStringWithKey(target string, key key) string {
	var result string

	for _, char := range []rune(target) {
		var processedCharNumber = processIntWithKey(int(char), key)
		result += string(processedCharNumber)
	}

	return result
}

func processIntWithKey(target int, key key) int {
	var result *big.Int = big.NewInt(int64(target))
	result.Exp(result, big.NewInt(int64(key.secondPiece)), nil)
	result.Mod(result, big.NewInt(int64(key.firstPiece)))

	return int(result.Int64())
}

func calculateN(p int, q int) int {
	return p * q
}

func calculateT(p int, q int) int {
	return (p - 1) * (q - 1)
}

func generateD(e int, t int) int {
	var result *big.Int = big.NewInt(int64(e))
	result.ModInverse(result, big.NewInt(int64(t)))

	return int(result.Int64())
}
