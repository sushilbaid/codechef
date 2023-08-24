package main

import (
	"bufio"
	"fmt"
	"os"
)

func breakstick() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n, x int
		fmt.Fscanln(reader, &n, &x)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s)
		odd1, odd2 := (n-x)%2, x%2
		if odd1 == odd2 {
			fmt.Println("yes")
		} else if odd1 == 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
