package main

import (
	"bufio"
	"fmt"
	"os"
)

func elections() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)
		r := "NOTA"
		switch {
		case a > 50:
			r = "A"
		case b > 50:
			r = "B"
		case c > 50:
			r = "c"
		}
		fmt.Println(r)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)

	}
}
