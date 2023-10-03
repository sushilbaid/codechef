package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, h int
		fmt.Fscanln(reader, &n, &h)
		s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		a := readNumbers(s)
		sort.Ints(a)
		r := 0
		for i := n - 1; i >= 0; i-- {
			if h > a[i] {
				h -= a[i]
			} else {
				r = a[i]
				break
			}
		}
		fmt.Println(r)
	}
}
