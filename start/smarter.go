package main

import (
	"bufio"
	"fmt"
	"os"
)

func smarter() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	// count = 1
	for t := 0; t < count; t++ {
		var l, v1, v2 int
		fmt.Fscanln(reader, &l, &v1, &v2)
		//s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		t1 := (l + v1 - 1) / v1
		t2 := (l + v2 - 1) / v2
		if t1 <= t2 {
			fmt.Println(-1)
		} else {
			fmt.Println(t1 - t2 - 1)
		}
	}
}
