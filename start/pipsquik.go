package main

import (
	"bufio"
	"fmt"
	"os"
)

func pipsquik() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, h, y1, y2, l int
		fmt.Fscanln(reader, &n, &h, &y1, &y2, &l)
		b, stopped := 0, false
		for i := 0; i < n; i++ {
			var t1, x1 int
			fmt.Fscanln(reader, &t1, &x1)
			if stopped {
				continue
			}
			switch t1 {
			case 1:
				if h-y1 > x1 && l <= 1 {
					stopped = true
					break
				} else if h-y1 > x1 {
					l--
				}
				b++
			case 2:
				if y2 < x1 && l <= 1 {
					stopped = true
					break
				} else if y2 < x1 {
					l--
				}
				b++
			}
		}
		fmt.Println(b)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)
	}
}
