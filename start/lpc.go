package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lpc() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		//a := readNumbers(s)
		k, _ := reader.ReadString('\n')
		k = strings.TrimSpace(k)
		min := 51
		for i := 0; i < n-m+1; i++ {
			sum := 0
			for j := 0; j < m; j++ {
				r1, r2 := int(s[i+j]), int(k[j])
				d := abs(r2 - r1)
				d = minint(d, 10-d)
				sum += d
			}
			if sum < min {
				min = sum
			}
			if sum == 0 {
				break
			}
		}
		fmt.Println(min)
	}
}

// abs return absolute value of a
func abs(a int) int {
	result := a
	if a < 0 {
		result = -a
	}
	return result
}

func minint(a, b int) int {
	result := a
	if b < a {
		result = b
	}
	return result
}
