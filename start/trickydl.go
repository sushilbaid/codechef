package main

import (
	"bufio"
	"fmt"
	"os"
)

func trickydl() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var a int
		fmt.Fscanln(reader, &a)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		// a := readNumbers(s)
		d1, d2 := 1, 1
		s1, s2, give, profit := 0, 0, 1, 0
		i := 0
		for {
			s1 += a
			s2 += give
			give *= 2
			if s2 > s1 {
				break
			}
			i++
			if s1-s2 > profit {
				profit = s1 - s2
				d2 = i
			}
		}
		d1 = i
		fmt.Println(d1, d2)
	}
}
