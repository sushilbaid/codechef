package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func mexsplit() {
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
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		// common mex can be zero or 1
		// 1 when number 0 is part of subarray. 1 when number>0 is part of array
		nonzeros := 0
		for _, v := range a {
			if v > 0 {
				nonzeros++
			}
		}
		r := maxint(n-nonzeros, nonzeros)
		fmt.Println(r)
	}
}
