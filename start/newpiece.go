package main

import (
	"bufio"
	"fmt"
	"os"
)

func newpiece() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var a, b, p, q int
		fmt.Fscanln(reader, &a, &b, &p, &q)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)
		c1, c2 := (a+b)%2 == 0, (p+q)%2 == 0
		if a == p && b == q {
			fmt.Println(0)
		} else if c1 != c2 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}
