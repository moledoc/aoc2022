package main

import (
	"fmt"
	"os"

	"github.com/moledoc/aoc2022/read"
)

const fname = "input.txt"

var input, err = read.File(fname)

func fst(distChar int) {
	line := input[0]
	l := len(line)
	var start int
	for i := 0; i < l; i++ {
		if i+4 < l {
			c := make(map[byte]bool)
			for j := i; j < i+distChar; j++ {
				c[line[j]] = true
			}
			if len(c) == distChar {
				start = i + distChar
				break
			}
		}
	}
	fmt.Println(start)
}

func main() {
	if err != nil {
		fmt.Printf("reading file '%v' failed\n", fname)
		os.Exit(1)
	}
	fst(4)
	fst(14)
}
