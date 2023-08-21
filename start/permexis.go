package main

import (
	"bufio"
	"codechef/utils"
	"fmt"
	"os"
	"sort"
)

func permixis() {
	abs := utils.NewMathLib().Abs
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
		sort.Ints(a)
		prev := a[0]
		result := "YES"
		for _, v := range a {
			if abs(v-prev) > 1 {
				result = "NO"
			}
			prev = v
		}
		fmt.Println(result)
	}
}
