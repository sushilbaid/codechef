package main

import (
	"bufio"
	"fmt"
	"os"
)

func tablet() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n, b int
		fmt.Fscanln(reader, &n, &b)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s)
		area := 0
		for i := 0; i < n; i++ {
			w, h, p := 0, 0, 0
			fmt.Fscanln(reader, &w, &h, &p)
			if p > b {
				continue
			}
			a := w * h
			if a > area {
				area = a
			}
		}
		if area > 0 {
			fmt.Println(area)
		} else {
			fmt.Println("no tablet")
		}
	}
}
