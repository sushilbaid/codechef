package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longseq() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		// var n int
		// fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		//a := readNumbers(s)
		zeros, ones := 0, 0
		for _, ch := range s {
			switch ch {
			case '1':
				ones++
			case '0':
				zeros++
			}
		}
		if ones == 1 || zeros == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
