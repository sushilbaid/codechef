package main

import (
	"bufio"
	"fmt"
	"os"
)

func spcp2() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var x, n int
		fmt.Fscanln(reader, &x, &n)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		x2 := (n + 99) / 100
		if x2 > x {
			fmt.Println(x2 - x)
		} else {
			fmt.Println(0)
		}
	}
}
