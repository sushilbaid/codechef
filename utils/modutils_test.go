package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type modusTesting struct {
	base  int
	modus int
	m     ModusLib
}

func newModusLibTesting() *modusTesting {
	base := 1000 * 1000 * 1000
	return &modusTesting{
		base:  base,
		modus: base + 7,
		m:     NewModusLib(base + 7),
	}
}

func TestAdd(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a, b int }{
		{4, 1, 3},
		{4, 3, 1},
		{5, mt.modus, 5},
		{4, 7 * mt.modus, mt.modus + 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=%d+%d", tt.result, tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Add(tt.a, tt.b))
		})
	}
}

func TestSub(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a, b int }{
		{2, 3, 1},
		{mt.modus - 2, 1, 3},
		{4, mt.modus + 4, mt.modus},
		{mt.modus - 4, mt.modus, mt.modus + 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=%d-%d", tt.result, tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Sub(tt.a, tt.b))
		})
	}
}

func TestMul(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a, b int }{
		{3, 1, 3},
		{3, 3, 1},
		{0, mt.modus, 5},
		{20, 7*mt.modus + 5, mt.modus + 4},
		{(1000 * 1000 * 1011) % mt.modus, 1000 * 1000, 1011},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=%d*%d", tt.result, tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Mul(tt.a, tt.b))
		})
	}
}

func TestInverse(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a int }{
		{1, 1},
		{333333336, 3},
		{500000004, mt.modus + 2},
	}
	for _, tt := range tests {
		a := assert.New(t)
		t.Run(fmt.Sprintf("%d=1/%d", tt.result, tt.a), func(t *testing.T) {
			a.Equal(tt.result, m.Inverse(tt.a))
		})
	}
}

func TestExp(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a, b int }{
		{1, 1, 3},
		{3, 3, 1},
		{243, 3, 5},
		{14348907, 3, 15},
		{486784380, 3, 20},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=%d^%d", tt.result, tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Exp(tt.a, tt.b))
		})
	}
}

func TestFact(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, a int }{
		{1, 1},
		{6, 3},
		{120, 5},
		{3628800, 10},
		{146326063, 20},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=factorial(%d)", tt.result, tt.a), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Fact(tt.a))
		})
	}
}

func TestFacts(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct {
		result []int
		a      int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 6}, 3},
		{[]int{1, 2, 6, 24, 120}, 5},
		{[]int{1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}, 10},
		{[]int{1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600, 227020758, 178290591, 674358851, 789741546, 425606191, 660911389, 557316307, 146326063}, 20},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("facts(%d)", tt.a), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Facts(tt.a)[1:])
		})
	}
}

func TestFibo(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct {
		result []int
		a      int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 3},
		{[]int{1, 1, 2, 3, 5}, 5},
		{[]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("fibonacci(%d)", tt.a), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Fibo(tt.a)[1:])
		})
		t.Run("fibonacci(50)", func(t *testing.T) {
			a := assert.New(t)
			// validate against last 5 in series
			expectedResult := []int{836311896, 971215059, 807526948, 778742000, 586268941}
			a.Equal(expectedResult, m.Fibo(50)[46:])
		})
	}
}

func TestNck(t *testing.T) {
	mt := newModusLibTesting()
	m := mt.m
	tests := []struct{ result, n, k int }{
		{1, 1, 1},
		{6, 4, 2},
		{0, mt.modus, 2},
		{4950, mt.modus + 100, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Nck(%d,%d)", tt.n, tt.k), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Nck(tt.n, tt.k))
		})
	}
}
