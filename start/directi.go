package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func directi() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		a := make([]string, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadString('\n')
			a[i] = strings.TrimSpace(s)
		}
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)
		b := make([]string, n)
		prev := ""
		for i := n - 1; i >= 0; i-- {
			tokens := strings.SplitN(a[i], " ", 2)
			turn := ""
			if prev == "" {
				turn = "Begin"
			} else if prev == "Right" {
				turn = "Left"
			} else {
				turn = "Right"
			}
			prev = tokens[0]
			b[n-1-i] = turn + " " + tokens[1]
		}
		for _, v := range b {
			fmt.Println(v)
		}
	}
}
