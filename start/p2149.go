package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func p2149() {
	minInt := utils.NewMathLib().MinInt
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var a, b, x int
		fmt.Fscanln(reader, &a, &b, &x)
		a1, a2 := a*b, x*x
		shorterEdge := minInt(a, b)
		var cost int
		if a1 <= a2 {
			cost = 0
		} else if a2 >= shorterEdge {
			cost = 1
		} else {
			cost = 2
		}
		fmt.Println(cost)
	}
}
