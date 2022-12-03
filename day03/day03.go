package main

import (
	"fmt"
	"os"

	"github.com/moledoc/aoc2022/read"
)

const fname = "input.txt"

var input, err = read.File(fname)

func fst() {
	result := 0
	for _, line := range input {
		contains := make(map[rune]int)
		l := len(line)
		frst := line[0 : l/2]
		scnd := line[l/2 : l]
		for _, c := range frst {
			contains[c] = 1
		}
		for _, c := range scnd {
			if count, ok := contains[c]; ok && count == 1 {
				if ic := int(c); 65 <= ic && ic <= 90 {
					result += ic - 38
				} else if 97 <= ic && ic <= 122 {
					result += ic - 96
				}
				contains[c] -= 1
			}
		}
	}
	fmt.Println(result)
}

func snd() {
	result := 0
	for i := 0; i < len(input); i += 3 {
		elf1 := input[i]
		elf2 := input[i+1]
		elf3 := input[i+2]
		contains := make(map[rune]int)
		for _, c := range elf1 {
			contains[c] = 1
		}
		for _, c := range elf2 {
			if contains[c] == 1 {
				contains[c] = 2
			}
		}
		for _, c := range elf3 {
			if count, ok := contains[c]; ok && count == 2 {
				if ic := int(c); 65 <= ic && ic <= 90 {
					result += ic - 38
				} else if 97 <= ic && ic <= 122 {
					result += ic - 96
				}
				contains[c] -= 1
			}
		}
	}
	fmt.Println(result)
}

func main() {
	if err != nil {
		fmt.Printf("reading file '%v' failed\n", fname)
		os.Exit(1)
	}
	fst()
	snd()
}
