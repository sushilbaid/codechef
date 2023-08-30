package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encmsg() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		var n int
		fmt.Fscanln(reader, &n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		//a := readNumbers(s)
		es := strings.Builder{}
		write := func(ch byte) {
			mapped := 'z' - (ch - 'a')
			es.WriteRune(rune(mapped))
		}
		l := len(s)
		for i := 0; i < l-1; i += 2 {
			write(s[i+1])
			write(s[i])
		}
		if l%2 == 1 {
			write(s[l-1])
		}
		fmt.Println(es.String())
	}
}
