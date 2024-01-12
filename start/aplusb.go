package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
	"sort"
)

func aplusb() {
	stringify := utils.NewStringLib().Stringify
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s1, _ := reader.ReadString('\n')
		s2, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		a := readNumbers(s1)
		b := readNumbers(s2)
		sort.Ints(a)
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})
		x := a[0] + b[0]
		r := true
		for i, v := range a[1:] {
			if v+b[i+1] != x {
				r = false
				break
			}
		}
		if r {
			fmt.Println(stringify(a))
			fmt.Println(stringify(b))
		} else {
			fmt.Println(-1)
		}
	}
}
