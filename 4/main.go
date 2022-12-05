package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cleaningRange struct {
	lowerBound, upperBound int
}

type elfPair struct {
	range1, range2 cleaningRange
}

func main() {
	pairs := load()
	fmt.Println(findFullyContained(pairs))
	fmt.Println(findAnyOverlap(pairs))
}

func findAnyOverlap(pairs []elfPair) int {
	total := 0
	for _, p := range pairs {
		if fullContain(p) ||
			p.range1.lowerBound >= p.range2.lowerBound && p.range1.lowerBound <= p.range2.upperBound ||
			p.range2.lowerBound >= p.range1.lowerBound && p.range2.lowerBound <= p.range1.upperBound {
			total++
		}
	}

	return total
}

func findFullyContained(pairs []elfPair) int {
	total := 0
	for _, p := range pairs {
		if fullContain(p) {
			total++
		}
	}

	return total
}

func fullContain(p elfPair) bool {
	if p.range1.lowerBound <= p.range2.lowerBound && p.range1.upperBound >= p.range2.upperBound {
		return true
	}
	if p.range2.lowerBound <= p.range1.lowerBound && p.range2.upperBound >= p.range1.upperBound {
		return true
	}
	return false
}

func load() []elfPair {
	infile, _ := os.Open("input")
	s := bufio.NewScanner(infile)

	pairs := []elfPair{}
	ranges := []cleaningRange{}

	for s.Scan() {
		sections := strings.Split(s.Text(), ",")
		for _, section := range sections {
			rangeBounds := strings.Split(section, "-")
			lowerBound, _ := strconv.Atoi(rangeBounds[0])
			upperBound, _ := strconv.Atoi(rangeBounds[1])
			ranges = append(ranges, cleaningRange{lowerBound, upperBound})
			if len(ranges) > 1 {
				pairs = append(pairs, elfPair{ranges[0], ranges[1]})
				ranges = []cleaningRange{}
			}
		}
	}

	return pairs
}
