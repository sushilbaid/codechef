package main

import (
	"bufio"
	"fmt"
	"os"
)

func hills() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, u, d int
		fmt.Fscanln(reader, &n, &u, &d)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		h := readNumbers(s)
		if len(h) == 1 {
			fmt.Println(1)
			continue
		}
		r, parashute := 1, false
		for i := 1; i < len(h); i++ {
			h1, h2 := h[i-1], h[i]
			if h2 > h1+u {
				break
			} else if h1 > h2+d {
				if !parashute {
					parashute = true
				} else {
					break
				}
			}
			r = i + 1
		}
		fmt.Println(r)
	}
}
