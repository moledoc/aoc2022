package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/moledoc/aoc2022/read"
)

const fname = "input.txt"

var input, err = read.File(fname)

func fst() {
	result := 0
	for _, line := range input {
		es := strings.Split(line, ",")

		r1 := strings.Split(es[0], "-")
		r1l, _ := strconv.Atoi(r1[0])
		r1u, _ := strconv.Atoi(r1[1])

		r2 := strings.Split(es[1], "-")
		r2l, _ := strconv.Atoi(r2[0])
		r2u, _ := strconv.Atoi(r2[1])

		if (r1l <= r2l && r2u <= r1u) || (r2l <= r1l && r1u <= r2u) {
			result += 1
		}
	}
	fmt.Println(result)
}

func snd() {
	result := 0
	for _, line := range input {
		es := strings.Split(line, ",")

		r1 := strings.Split(es[0], "-")
		r1l, _ := strconv.Atoi(r1[0])
		r1u, _ := strconv.Atoi(r1[1])

		r2 := strings.Split(es[1], "-")
		r2l, _ := strconv.Atoi(r2[0])
		r2u, _ := strconv.Atoi(r2[1])

		if (r1l <= r2l && r2l <= r1u) || (r2l <= r1l && r1l <= r2u) {
			result += 1
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
