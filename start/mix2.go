package main

import (
	"bufio"
	"fmt"
	"os"
)

func mix2() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	// count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s)
		r := 0
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				r += a[i] * a[j]
			}
		}
		fmt.Println(r)
	}
}
