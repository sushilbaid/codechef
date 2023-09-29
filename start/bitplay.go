package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bitplay() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		// a := readNumbers(s)
		r := 1
		for i := 0; i < n-2; i += 2 {
			s1, s2, s3 := s[i], s[i+1], s[i+2]
			s1, s2, s3 = s1-'0', s2-'0', s3-'0'
			options := 0
			if (s1 & s2) == s3 {
				options++
			}
			if (s1 | s2) == s3 {
				options++
			}
			if (s1 ^ s2) == s3 {
				options++
			}
			fmt.Println(s1, s2, s3, options)
			const MODUS = 1000*1000*1000 + 7
			r = (r * options) % MODUS
			if r == 0 {
				break
			}
		}
		fmt.Println(r)
	}
}
