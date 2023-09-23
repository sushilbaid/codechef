package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func kdeli() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, k, l int
		fmt.Fscanln(reader, &n, &k, &l)
		s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		sort.Ints(a)
		i, sum := n-l, a[n-l]
		for {
			i -= k
			if i < 0 {
				break
			}
			sum += a[i]
		}
		fmt.Println(sum)
	}
}
