package main

import (
	"bufio"
	"fmt"
	"os"
)

func spcp3() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		if a%b == 0 {
			fmt.Println(0)
			continue
		}
		x, y := a, b
		r1, r2 := 0, 0
		for x%y != 0 {
			x++
			y--
			r1++
		}
		x, y = a, b
		if x > y {
			for x >= y && x%y != 0 {
				x--
				y++
				r2++
			}
			if x%y == 0 && r2 < r1 {
				r1 = r2
			}
		}
		fmt.Println(r1)
	}
}
