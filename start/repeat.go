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
		var n, k, s int
		fmt.Fscanln(reader, &n, &k, &s)
		//s, _ := reader.ReadString('\n')
		//s = strings.TrimSpace(s)
		//a := readNumbers(s)
		x := (s - n*n) / (k - 1)
		fmt.Println(x)
	}
}
