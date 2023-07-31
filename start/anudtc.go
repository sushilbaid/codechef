package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
		x := 360%n == 0
		y := n <= 360
		z := n <= 26
		result := strings.Builder{}
		for _, v := range []bool{x, y, z} {
			if v {
				result.WriteString(" y")
			} else {
				result.WriteString(" n")
			}
		}
		fmt.Println(strings.TrimSpace(result.String()))
	}
}
