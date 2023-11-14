package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	fmt.Fscanln(reader, &count)
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		//s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		//a := readNumbers(s)
		if n == 1 || n%2 == 0 {
			fmt.Println("Bob")
		} else {
			fmt.Println("Alice")
		}
	}
}
