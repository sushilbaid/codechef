package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func eventual() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		// a := readNumbers(s)
		m := map[rune]int{}
		for _, v := range s {
			m[v]++
		}
		r := true
		for _, x := range m {
			if x%2 == 1 {
				r = false
				break
			}
		}
		if r {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
