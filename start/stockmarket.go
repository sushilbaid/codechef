package main

import (
	"bufio"
	"fmt"
	"os"
)

func stockmarket() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s)
		p, min := 0, 100
		for _, v := range a {
			p += v
			if min > v {
				min = v
			}
		}
		p = p - min
		fmt.Println(p)
	}
}
