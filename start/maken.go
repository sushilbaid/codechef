package main

import (
	"bufio"
	"fmt"
	"os"
)

func maken() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	// count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		//s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		a := make([][][]int, n)
		for i := 0; i < n; i++ {
			r := [][]int{}
			for j := 0; j < i; j++ {
				for _, v := range a[j] {
					l := len(v)
					y := i - j
					if v[l-1] > y {
						continue
					}
					x := make([]int, len(v))
					copy(x, v)
					x = append(x, y)
					r = append(r, x)
				}
			}
			r = append(r, []int{i + 1})
			a[i] = r
			//fmt.Println(a[i])
		}
		for _, v := range a[n-1] {
			fmt.Println(v)
		}
	}
}
