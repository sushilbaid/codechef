package main

import (
	"bufio"
	"fmt"
	"os"
)

func muffins3() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		// s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		// a := readNumbers(s)
		r := 2
		if n > 2 {
			r = n/2 + 1
		}
		fmt.Println(r)
	}
}
