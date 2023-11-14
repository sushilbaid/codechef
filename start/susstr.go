package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func susstr() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		//a := readNumbers(s)
		r := ""
		alice := true
		i, j := 0, len(s)-1
		for i <= j {
			if alice {
				ch := s[i]
				switch ch {
				case '1':
					r = r + "1"
				case '0':
					r = "0" + r
				}
				i++
				alice = false
			} else {
				ch := s[j]
				switch ch {
				case '1':
					r = "1" + r
				case '0':
					r = r + "0"
				}
				j--
				alice = true
			}
		}
		fmt.Println(r)
	}
}
