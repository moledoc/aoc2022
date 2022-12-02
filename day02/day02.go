package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/moledoc/aoc2022/read"
)

const fname = "input.txt"

var (
	input, err = read.File(fname)
	rps        = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
)

func fst() {
	score := 0
	for _, line := range input {
		elems := strings.Split(line, " ")
		you := rps[elems[1]]
		score += you
		opponent := rps[elems[0]]
		if diff := you - opponent; diff == 1 || diff == -2 { // -2 is rock beats scissors
			score += 6
		} else if diff := you - opponent; diff == 0 {
			score += 3
		}
	}
	fmt.Println(score)
}

func snd() {
	score := 0
	for _, line := range input {
		elems := strings.Split(line, " ")
		roundResult := (rps[elems[1]] - 1) * 3 // lose, draw, win
		score += roundResult
		opponent := rps[elems[0]]
		if roundResult == 6 { // win
			play := (opponent + 1) % 4
			if play == 0 {
				play = 1
			}
			score += play
		} else if roundResult == 3 {
			score += opponent
		} else if roundResult == 0 {
			play := opponent - 1%4
			if play == 0 {
				play = 3
			}
			score += play
		}
	}
	fmt.Println(score)
}

func main() {
	if err != nil {
		fmt.Printf("reading file '%v' failed\n", fname)
		os.Exit(1)
	}
	fst()
	snd()
}
