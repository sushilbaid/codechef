package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
)

func plprocess() {
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
		if n == 1 {
			fmt.Println(a[0])
			continue
		}

		i1, i2, sum1, sum2 := 0, n-1, a[0], a[n-1]
		for {
			if i2-i1 <= 1 {
				break
			}
			if sum1+a[i1+1] < sum2+a[i2-1] {
				i1, sum1 = i1+1, sum1+a[i1+1]
			} else {
				i2, sum2 = i2-1, sum2+a[i2-1]
			}
		}
		result := maxint(sum1, sum2)
		fmt.Println(result)
	}
}
