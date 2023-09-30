package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func removecards() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		sort.Ints(a)
		prev, max, c := 0, 0, 1
		for _, v := range a {
			if v != prev {
				if c > max {
					max = c
				}
				prev, c = v, 1
			} else {
				c++
			}
		}
		if c > max {
			max = c
		}
		fmt.Println(n - max)
	}
}
