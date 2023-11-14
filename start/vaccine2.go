package main

import (
	"bufio"
	"fmt"
	"os"
)

func vaccine2() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, d int
		fmt.Fscanln(reader, &n, &d)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s)
		rc := 0
		for _, v := range a {
			if v >= 90 || v <= 9 {
				rc++
			}
		}
		c1, c2 := (rc+d-1)/d, (n-rc+d-1)/d
		r := c1 + c2
		fmt.Println(r)
	}
}
