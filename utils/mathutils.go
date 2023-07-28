package utils

import (
	"math"
	"time"
)

type mathLib struct {
}

type MathLib interface {
	Abs(a int) int
	MinInt(a, b int) int
	MaxInt(a, b int) int
	FindMax(a []int) int
	Gcd(a, b int) int
	MaskMsb(n int) int
	GeneratePrimes(n int) ([]int, time.Duration)
}

func NewMathLib() MathLib {
	return &mathLib{}
}

// abs return absolute value of a
func (*mathLib) Abs(a int) int {
	result := a
	if a < 0 {
		result = -a
	}
	return result
}

func (*mathLib) MinInt(a, b int) int {
	result := a
	if b < a {
		result = b
	}
	return result
}

func (*mathLib) MaxInt(a, b int) int {
	result := a
	if b > a {
		result = b
	}
	return result
}

// FindMax return the max int in the given slice
func (*mathLib) FindMax(a []int) int {
	max := math.MinInt
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

// MaskMsb returns mask with most significant bit of given n on
func (*mathLib) MaskMsb(n int) int {
	if n == 0 {
		return 0
	}

	m := 1
	for n > 0 {
		n >>= 1
		if n == 0 {
			break
		}
		m <<= 1
	}
	return m
}

// gcd computes greatest common divisor of given numbers
func (*mathLib) Gcd(a int, b int) int {
	divisor, remainder := a, b
	for remainder != 0 {
		divisor, remainder = remainder, divisor%remainder
	}

	return divisor
}

// GeneratePrimes computes primes <= n and return them including time duration taken to compute them
func (*mathLib) GeneratePrimes(n int) ([]int, time.Duration) {
	t := time.Now()
	a := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		a[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if !a[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			a[j] = false
		}
	}
	result := []int{}
	for i := 2; i <= n; i++ {
		if a[i] {
			result = append(result, i)
		}
	}
	taken := time.Since(t)

	return result, taken
}
