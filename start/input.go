package main

import (
	"strconv"
	"strings"
)

func readNumbers(s string) []int {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	numbers := []int{}
	for _, str := range strs {
		n, _ := strconv.Atoi(str)
		numbers = append(numbers, n)
	}
	return numbers
}
