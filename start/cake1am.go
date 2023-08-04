package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func cake1am() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var x1, y1, x2, y2 int
		fmt.Fscanln(reader, &x1, &y1, &x2, &y2)
		var x3, y3, x4, y4 int
		fmt.Fscanln(reader, &x3, &y3, &x4, &y4)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		// a := readNumbers(s)
		a1, a2 := (x2-x1)*(y2-y1), (x4-x3)*(y4-y3)
		var r int
		// non overlapping
		if (x4 <= x1 || x2 <= x3) ||
			(y4 <= y1 || y2 <= y3) {
			r = a1 + a2
		} else if x3 >= x1 && x4 <= x2 &&
			y3 >= y1 && y4 <= y2 {
			// full overlap
			r = a1
		} else if x1 >= x3 && x2 <= x4 &&
			y1 >= y3 && y2 <= y4 {
			// full overlap
			r = a2
		} else {
			// partial overlap
			x := []int{x1, x2, x3, x4}
			sort.Ints(x)
			y := []int{y1, y2, y3, y4}
			sort.Ints(y)
			overlap := (x[2] - x[1]) * (y[2] - y[1])
			r = a1 + a2 - overlap
		}
		fmt.Println(r)
	}
}
