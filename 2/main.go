package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type move string
type scoreMap map[move]map[move]int

var outcomes = map[move]map[move]int{
	"A": {
		"X": 4,
		"Y": 8,
		"Z": 3,
	},
	"B": {
		"X": 1,
		"Y": 5,
		"Z": 9,
	},
	"C": {
		"X": 7,
		"Y": 2,
		"Z": 6,
	},
}

var outcomes2 = map[move]map[move]int{
	"A": {
		"X": 3,
		"Y": 4,
		"Z": 8,
	},
	"B": {
		"X": 1,
		"Y": 5,
		"Z": 9,
	},
	"C": {
		"X": 2,
		"Y": 6,
		"Z": 7,
	},
}

func main() {
	fmt.Println(calcScoreWithOutcomes(outcomes))
	fmt.Println(calcScoreWithOutcomes(outcomes2))
}

func calcScoreWithOutcomes(scores scoreMap) int {
	infile, _ := os.Open("input")
	defer infile.Close()
	s := bufio.NewScanner(infile)
	acc := 0
	for s.Scan() {
		t := s.Text()
		moves := strings.Split(t, " ")
		them, me := moves[0], moves[1]
		score := scores[move(them)][move(me)]
		acc += score
	}
	return acc
}
