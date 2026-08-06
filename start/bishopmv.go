package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)
	for i := 0; i < tc; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		if ((x1 + y1) % 2) != ((x2 + y2) % 2) {
			fmt.Fprintln(out, "-1")
		} else {
			if (x1+y1 == x2+y2) ||
				(x1-y1 == x2-y2) {
				fmt.Fprintln(out, "1")
			} else {
				fmt.Fprintln(out, "2")
			}
		}
	}
}
