package main

import (
	"bufio"
	"fmt"
	"os"
)

func candy123() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)
		var result string
		x, y, c1, c2 := 1, 2, 0, 0
		for {
			if c1+x > a {
				result = "Bob"
				break
			}
			c1 += x
			x += 2
			if c2+y > b {
				result = "Limak"
				break
			}
			c2 += y
			y += 2
		}
		fmt.Println(result)
	}
}
