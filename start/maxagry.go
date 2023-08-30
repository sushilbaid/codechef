package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func maxagry() {
	minint := utils.NewMathLib().MinInt
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, k int
		fmt.Fscanln(reader, &n, &k)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s)

		// max swap required to reverse the sort order is n/2
		c := minint(n/2, k)
		// each swap will generate 2 numbers in the arithmetic series n-1, n-2, ..,1,0
		// sum of arithmetic series = n*(2a-(n-1)d)/2
		sum := c * (2*(n-1) - (2*c - 1))
		fmt.Println(sum)
	}
}
