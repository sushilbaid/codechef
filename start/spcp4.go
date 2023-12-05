package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func spcp4() {
	abs := utils.NewMathLib().Abs
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, x, k int
		fmt.Fscanln(reader, &n, &x, &k)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		y := n - x
		x2 := x % k
		y2 := y % k
		r := abs(x2 - y2)
		fmt.Println(r)
	}
}
