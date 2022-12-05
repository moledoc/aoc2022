package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/moledoc/aoc2022/read"
)

const fname = "input.txt"

var input, err = read.File(fname)

func fst() {
	// stacks
	cont := 0
	stacks := make(map[int][]byte)
	for i, line := range input {
		if line == "" {
			cont = i + 1
			break
		}
		col := 0
		for j := 0; j < len(line); j += 4 {
			col++
			if line[j] != '[' {
				continue
			}
			stacks[col] = append(stacks[col], line[j+1])
		}
	}
	// commands
	for i := cont; i < len(input); i++ {
		line := input[i]
		re := regexp.MustCompile(`move (.*) from (.) to (.)`)
		match := re.FindStringSubmatch(line)
		howmany, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		l := len(stacks[to])
		n := make([]byte, l+howmany)
		for i := howmany - 1; i >= 0; i-- {
			n[howmany-i-1] = stacks[from][i]
		}
		for i := 0; i < l; i++ {
			n[howmany+i] = stacks[to][i]
		}
		stacks[to] = n
		stacks[from] = stacks[from][howmany:]
	}
	var result []byte
	for i := 1; i <= len(stacks); i++ {
		if len(stacks[i]) > 0 {
			result = append(result, stacks[i][0])
		}
	}
	fmt.Println(string(result))
}

func snd() {
	// stacks
	cont := 0
	stacks := make(map[int][]byte)
	for i, line := range input {
		if line == "" {
			cont = i + 1
			break
		}
		col := 0
		for j := 0; j < len(line); j += 4 {
			col++
			if line[j] != '[' {
				continue
			}
			stacks[col] = append(stacks[col], line[j+1])
		}
	}
	// commands
	for i := cont; i < len(input); i++ {
		line := input[i]
		re := regexp.MustCompile(`move (.*) from (.) to (.)`)
		match := re.FindStringSubmatch(line)
		howmany, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])

		l := len(stacks[to])
		n := make([]byte, l+howmany)
		for i := 0; i < howmany; i++ {
			n[i] = stacks[from][i]
		}
		for i := 0; i < l; i++ {
			n[howmany+i] = stacks[to][i]
		}
		stacks[to] = n
		stacks[from] = stacks[from][howmany:]
	}
	var result []byte
	for i := 1; i <= len(stacks); i++ {
		if len(stacks[i]) > 0 {
			result = append(result, stacks[i][0])
		}
	}
	fmt.Println(string(result))
}

func main() {
	if err != nil {
		fmt.Printf("reading file '%v' failed\n", fname)
		os.Exit(1)
	}
	fst()
	snd()
}
