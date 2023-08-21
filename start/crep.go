package main

import (
	"bufio"
	"fmt"
	"os"
)

func crep() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, k int
		fmt.Fscanln(reader, &n, &k)
		s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		m := map[int]int{} // map of number to repetitions
		for _, v := range a {
			m[v]++
		}
		sum, found := 0, false
		for number, v := range m {
			if v == k {
				sum += number
				found = true
			}
		}
		if !found {
			fmt.Println(-1)
		} else {
			fmt.Println(sum)
		}
	}
}
