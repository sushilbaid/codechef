package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func main() {
	maxint := utils.NewMathLib().MaxInt
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s)
		m := map[int]bool{}
		max, c := 0, 0
		for _, v := range a {
			if m[v] {
				m[v] = false
				c--
			} else {
				m[v] = true
				c++
			}
			max = maxint(max, c)
		}
		fmt.Println(max)
	}
}
