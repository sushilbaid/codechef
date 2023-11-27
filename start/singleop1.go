package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func singleop1() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		// trim 0s on left
		s = strings.TrimLeft(s, "0")
		y := 0
		// all ones until a zero need to be shifted so that first zero becomes one after xor
		// e.g. 1111011 needs to be shifted by 4 bits. that is y = 4. in other words, divide x by 2^4 and y=4.
		for _, v := range s {
			if v == '0' {
				break
			}
			y++
		}
		fmt.Println(y)
	}
}
