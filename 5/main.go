package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile, _ := os.Open("input")
	s := bufio.NewScanner(infile)

	// First chunk is the starting state of the stacks
	numStacks := 0
	stacks := [][]string{}
	for s.Scan() {
		t := s.Text()
		if numStacks == 0 {
			numStacks = (len(t) + 1) / 4
			for i := 0; i < numStacks; i++ {
				stacks = append(stacks, []string{})
			}
		}
		// A blank line indicates that we're done building the starting state, and the rest of the file is moves.
		if t == "" {
			break
		}

		for i := 0; i < numStacks; i++ {
			stackPosition := string(t[i*4+1])
			if _, err := strconv.Atoi(stackPosition); err == nil {
				continue
			}
			if stackPosition != " " {
				stacks[i] = append(stacks[i], stackPosition)
			}
		}
	}

	// Reverse all the slices since we built them from the top down, but want them to describe bottom up.
	for i, s := range stacks {
		stacks[i] = reverseSlice(s)
		fmt.Println(i+1, " ", stacks[i])
	}

	// Next, do the mutations on the stacks
	for s.Scan() {
		t := strings.Split(s.Text(), " ")
		fmt.Println(t[1], t[3], t[5])
		move, _ := strconv.Atoi(t[1])
		from, _ := strconv.Atoi(t[3])
		to, _ := strconv.Atoi(t[5])
		from--
		to--
		// singleMove(stacks, move, from, to)
		multiMove(stacks, move, from, to)
	}

	// Print the top of each stack.
	for _, s := range stacks {
		fmt.Println(s[len(s)-1])
	}
}

func multiMove(stacks [][]string, move, from, to int) {
	if len(stacks[from]) >= move {
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-move:]...)
		stacks[from] = (stacks[from][:len(stacks[from])-move])
	}
}

func singleMove(stacks [][]string, move, from, to int) {
	for i := 0; i < move; i++ {
		if len(stacks[from]) > 0 {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
}

func reverseSlice(s []string) []string {
	reversed := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}
