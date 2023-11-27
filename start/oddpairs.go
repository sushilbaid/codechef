package main

import (
	"bufio"
	"fmt"
	"os"
)

func oddpairs() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		odd, even := (n+1)/2, n/2
		r := odd * even * 2
		fmt.Println(r)
	}
}
