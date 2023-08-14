package main

import (
	"bufio"
	"fmt"
	"os"
)

func twodiffpalin() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s1)
		if a%2 == 1 && b%2 == 1 {
			fmt.Println("no")
		} else if a == 1 || b == 1 {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}
