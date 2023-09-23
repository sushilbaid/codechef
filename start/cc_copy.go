package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type token struct {
	ch    rune
	index int
}

const codechef = "codechef"

func cc_copy() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	var count int
	fmt.Fscanln(reader, &count)
	//count = 1
	for t := 0; t < count; t++ {
		//var n, k, l int
		//fmt.Fscanln(reader, &n, &k, &l)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		// a := readNumbers(s)

		// validate max count of characters of possibility of a solution
		max := map[rune]int{
			'c': 6, 'o': 7, 'd': 7, 'e': 6, 'h': 7, 'f': 7,
		}
		m := map[rune]int{}
		for _, v := range s {
			m[v]++
		}
		possible := true
		for k, c := range m {
			if x, ok := max[k]; ok && c > x {
				possible = false
				break
			}
		}
		if !possible {
			fmt.Println(-1)
			continue
		}

		// identify mislocated characters
		tokens := make([]*token, 8)
		for i, v := range s {
			tokens[i] = &token{v, i}
		}

		locations := needChangeOfLocation(tokens)

		// 1. first try to swap among these
		i, j := 0, len(locations)-1
		for {
			if i >= j {
				break
			}
			swap(locations[i], locations[j])
			i++
			j--
		}
		locations = needChangeOfLocation(tokens)

		// 2. now search in full list
		for _, location := range locations {
			i := location.index
			for _, tk := range tokens {
				j := tk.index
				if i != j && rune(codechef[i]) != tk.ch && rune(codechef[j]) != location.ch {
					swap(location, tk)
					break
				}
			}
		}

		// build result string
		runes := make([]rune, 8)
		for _, tk := range tokens {
			runes[tk.index] = tk.ch
		}
		fmt.Println(string(runes))
	}
}

func swap(t1, t2 *token) {
	t1.ch, t2.ch = t2.ch, t1.ch
}

func needChangeOfLocation(tokens []*token) []*token {
	locations := []*token{}
	for _, tk := range tokens {
		i := tk.index
		if tk.ch == rune(codechef[i]) {
			// need change of location
			locations = append(locations, tk)
		}
	}
	return locations
}
