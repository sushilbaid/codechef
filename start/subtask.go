package main

import (
	"bufio"
	"fmt"
	"os"
)

func subtask() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, m, k int
		fmt.Fscanln(reader, &n, &m, &k)
		s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		x := 0
		for _, v := range a {
			if v == 0 {
				break
			}
			x++
		}
		r := 0
		switch {
		case x == n:
			r = 100
		case x > m:
			r = k
		}
		fmt.Println(r)
	}
}
