package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func chrisdecor() {
	minint := utils.NewMathLib().MinInt
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, x, y int
		fmt.Fscanln(reader, &n, &x, &y)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		l1 := minint(x, y/3)
		max := (x-l1)/2 + l1
		if n <= max {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
