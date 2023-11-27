package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func singleop1() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		//a := readNumbers(s)
		s = strings.TrimLeft(s, "0")
		y := 0
		for _, v := range s {
			if v == '0' {
				break
			}
			y++
		}
		fmt.Println(y)
	}
}
