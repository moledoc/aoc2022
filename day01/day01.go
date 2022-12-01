package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/moledoc/aoc2022/read"
)

const (
	fname = "input.txt"
)

var (
	input, err = read.File(fname)
)

func fst() {
	// naive solution
	most := 0
	count := 0
	for _, cal := range input {
		if cal == "" {
			if count > most {
				most = count
			}
			count = 0
			continue
		}
		c, err := strconv.Atoi(cal)
		if err != nil {
			fmt.Printf("unable to convert '%v' to int\n", cal)
			continue
		}
		count += c
	}
	fmt.Println(most)
}

func top3(cals *[]int, count int) {
	if count > (*cals)[2] {
		(*cals)[0] = (*cals)[1]
		(*cals)[1] = (*cals)[2]
		(*cals)[2] = count
	} else if count > (*cals)[1] {
		(*cals)[0] = (*cals)[1]
		(*cals)[1] = count
	} else if count > (*cals)[0] {
		(*cals)[0] = count
	}
}

func snd() {
	cals := make([]int, 3)
	count := 0
	for _, cal := range input {
		if cal == "" {
			top3(&cals, count)
			count = 0
			continue
		}
		c, err := strconv.Atoi(cal)
		if err != nil {
			fmt.Printf("unable to convert '%v' to int\n", cal)
			continue
		}
		count += c
	}
	top3(&cals, count)
	fmt.Println(cals[0] + cals[1] + cals[2])
}

func main() {
	if err != nil {
		fmt.Printf("reading file '%v' failed\n", fname)
		os.Exit(1)
	}
	fst()
	snd()
}
