package main

import (
	"bufio"
	"fmt"
	"os"
)

func diet() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, k int
		fmt.Fscanln(reader, &n, &k)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s)
		r, d := 0, 0
		for i, v := range a {
			if r+v < k {
				d = i + 1
				break
			}
			r += v - k
		}
		if d == 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no", d)
		}
	}
}
