package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// minint := utils.NewMathLib().MaxInt
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		// s, _ := reader.ReadString('\n')
		// s = strings.TrimSpace(s)
		// a := readNumbers(s)
	}
}
