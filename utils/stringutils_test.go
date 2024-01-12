package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		input          []int
		expectedResult string
	}{
		{nil, ""},
		{[]int{1, 2}, "1 2 "},
		{[]int{100, 200}, "100 200"},
	}
	stringify := NewStringLib().Stringify
	for _, test := range tests {
		actual := stringify(test.input)
		a.Equal(test.expectedResult, actual)
	}
}
