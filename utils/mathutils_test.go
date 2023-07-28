package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	m := NewMathLib()
	tests := []struct{ result, a int }{
		{1, 1},
		{3, -3},
		{100000, -100000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("abs(%d)", tt.a), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Abs(tt.a))
		})
	}
}

func TestMinInt(t *testing.T) {
	m := NewMathLib()
	tests := []struct{ result, a, b int }{
		{1, 1, 2},
		{-3, 1, -3},
		{0, 1, 0},
		{-100000, 100000, -100000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("minint(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.MinInt(tt.a, tt.b))
		})
	}
}

func TestMaxInt(t *testing.T) {
	m := NewMathLib()
	tests := []struct{ result, a, b int }{
		{2, 1, 2},
		{1, 1, -3},
		{1, 1, 0},
		{100000, 100000, -100000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("maxint(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.MaxInt(tt.a, tt.b))
		})
	}
}

func TestFindMax(t *testing.T) {
	m := NewMathLib()
	tests := []struct {
		result int
		a      []int
	}{
		{2, []int{1, 2, 2}},
		{3, []int{1, -3, 3, -2}},
		{10, []int{1, 4, 2, 10, 9}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("findmax(%d)", tt.result), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.FindMax(tt.a))
		})
	}
}

func TestMaskMsb(t *testing.T) {
	m := NewMathLib()
	tests := []struct{ result, a int }{
		{0, 0},
		{1, 1},
		{2, 3},
		{8, 15},
		{16, 16},
		{16, 20},
		{1024, 1048},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("maskmsb(%d)", tt.a), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.MaskMsb(tt.a))
		})
	}
}

func TestGcd(t *testing.T) {
	m := NewMathLib()
	tests := []struct{ result, a, b int }{
		{1, 1, 2},
		{2, 2, 6},
		{3, 6, 9},
		{6, 30, 42},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("gcd(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(tt.result, m.Gcd(tt.a, tt.b))
		})
	}
}

func TestGeneratePrimes(t *testing.T) {
	m := NewMathLib()
	tests := []struct {
		result []int
		n      int
	}{
		{[]int{2}, 2},
		{[]int{2, 3}, 3},
		{[]int{2, 3, 5, 7, 11}, 12},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}, 100},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("generateprimes(%d)", tt.n), func(t *testing.T) {
			a := assert.New(t)
			r, _ := m.GeneratePrimes(tt.n)
			a.Equal(tt.result, r)
		})
	}
}
