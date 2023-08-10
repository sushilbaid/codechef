package main

import (
	"bufio"
	"fmt"
	"os"
)

func cs2023_stk() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s1, _ := reader.ReadString('\n')
		s2, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s1)
		b := readNumbers(s2)
		c1, c2 := findMax(a), findMax(b)
		if c1 > c2 {
			fmt.Println("Om")
		} else if c1 < c2 {
			fmt.Println("Addy")
		} else {
			fmt.Println("Draw")
		}
	}
}

func findMax(a []int) int {
	c, c1 := 0, 0
	for _, v := range a {
		if v > 0 {
			c++
		} else {
			c = 0
		}
		if c > c1 {
			c1 = c
		}
	}
	return c1
}
