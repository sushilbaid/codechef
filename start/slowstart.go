package main

import (
	"bufio"
	"fmt"
	"os"
)

func slowstart() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	// count = 1
	for t := 0; t < count; t++ {
		var x, h int
		fmt.Fscanln(reader, &x, &h)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		// a := readNumbers(s)
		x2 := x / 2
		var r int
		if h <= 5*x2 {
			r = (h + x2 - 1) / x2
		} else {
			h -= 5 * x2
			r = 5 + (h+x-1)/x
		}
		fmt.Println(r)
	}
}
