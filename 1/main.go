package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	index, total int
}

func (e Elf) String() string {
	return fmt.Sprintf("Elf %d is carrying %d calories", e.index, e.total)
}

func main() {
	infile, _ := os.Open("input")
	defer infile.Close()

	perElf := calcPerElfTotal(infile)
	getTopNPackElves(1, perElf)
	getTopNPackElves(3, perElf)
}

func calcPerElfTotal(inp io.Reader) []Elf {
	ind := 0
	elves := make([]Elf, 1)

	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			elves = append(elves, Elf{ind, 0})
			ind++
			continue
		} else {
			amt, _ := strconv.Atoi(t)
			elves[ind].total += amt
		}
	}
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].total < elves[j].total
	})
	return elves
}

func getTopNPackElves(n int, elfSlice []Elf) {
	fmt.Printf("Top %d ----\n", n)
	total := 0
	for _, v := range elfSlice[len(elfSlice)-n:] {
		fmt.Println(v)
		total += v.total
	}
	fmt.Printf("total calories: %d\n", total)
}
