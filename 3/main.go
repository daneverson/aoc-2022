package main

import (
	"bufio"
	"fmt"
	"os"
)

type rucksack struct {
	compartment1, compartment2 []rune
	sharedItem                 rune
	sharedItemPriority         int
}

func read(s string) *rucksack {
	if len(s)%2 != 0 {
		fmt.Println("Expected string to be of even length")
		os.Exit(1)
	}

	r := rucksack{}
	s1, s2 := s[:len(s)/2], s[len(s)/2:]
	members := map[rune]bool{}
	foundShared := false
	for _, c := range s1 {
		r.compartment1 = append(r.compartment1, c)
		members[c] = true
	}
	for _, c := range s2 {
		r.compartment2 = append(r.compartment2, c)
		if _, ok := members[c]; ok {
			foundShared = true
			r.sharedItem = c
			if r.sharedItem < 91 {
				r.sharedItemPriority = int(r.sharedItem - 38)
			} else {
				r.sharedItemPriority = int(r.sharedItem - 96)
			}
		}
	}

	if !foundShared {
		fmt.Println("expected one item to be shared between compartments")
		os.Exit(2)
	}

	return &r
}

func main() {
	infile, _ := os.Open("input")
	defer infile.Close()

	rucksacks := []*rucksack{}
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		e1 := scanner.Text()
		scanner.Scan()
		e2 := scanner.Text()
		scanner.Scan()
		e3 := scanner.Text()

		_, pri := findCommon(e1, e2, e3)
		r1 := read(e1)
		r1.sharedItemPriority = pri
		rucksacks = append(rucksacks, r1)
	}
	sumOfShared := 0
	for _, r := range rucksacks {
		fmt.Println(string(r.sharedItem), r.sharedItem, r.sharedItemPriority)
		sumOfShared += r.sharedItemPriority
	}
	fmt.Println(sumOfShared)
}

func findCommon(s1, s2, s3 string) (rune, int) {
	fmt.Printf("*** checking:\n\t%s\n\t%s\n\t%s\n", s1, s2, s3)
	found := map[rune]bool{}
	common12 := map[rune]bool{}
	common23 := map[rune]bool{}
	var commonElement rune

	for _, c := range s1 {
		found[c] = true
	}
	for _, c := range s2 {
		if found[c] {
			common12[c] = true
		}
	}
	for _, c := range s3 {
		if common12[c] {
			common23[c] = true
			commonElement = c
		}
	}
	if len(common23) != 1 {
		fmt.Println("expected exactly one element common between group of 3")
		os.Exit(3)
	}

	pri := 0
	if commonElement < 91 {
		pri = int(commonElement - 38)
	} else {
		pri = int(commonElement - 96)
	}
	return commonElement, pri
}
