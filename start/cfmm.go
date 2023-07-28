package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
	"strings"
)

func cfmm() {
	minint := utils.NewMathLib().MinInt
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		m := map[rune]int{}
		for _, ch := range "codechef" {
			m[ch] = 0
		}
		for i := 0; i < n; i++ {
			s, _ := reader.ReadString('\n')
			s = strings.TrimSpace(s)
			for _, ch := range s {
				if _, ok := m[ch]; ok {
					m[ch]++
				}
			}
		}
		min := minint(m['c']/2, m['e']/2)
		min = minint(min, m['o'])
		min = minint(min, m['d'])
		min = minint(min, m['h'])
		min = minint(min, m['f'])
		fmt.Println(min)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		// a := readNumbers(s)
	}
}
