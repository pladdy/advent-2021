package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bitstring struct {
	bits  []int
	value int64
}

func newBitstring(s string) bitstring {
	bits := []int{}
	for _, c := range s {
		if string(c) == "1" {
			bits = append(bits, 1)
		} else {
			bits = append(bits, 0)
		}
	}

	value, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return bitstring{bits, value}
}

func part1() {
	file, err := os.Open("day-3-puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	bitCounters := []map[string]int{}

	// Store count of bits we read from the puzzle file in a slice of maps that
	// track how many 1's and 0's are in each bit position.
	for _ = range [12]int{} {
		bitCounters = append(bitCounters, map[string]int{})
	}

	// Parse bit string and increment bitCounters
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		for i, c := range line {
			bitCounters[i][string(c)]++
		}
	}

	// Calculate epsilon/gamma based on bitCounters
	epsilon_string := strings.Builder{}
	gamma_string := strings.Builder{}
	for _, bitCounter := range bitCounters {
		zeros := bitCounter["0"]
		ones := bitCounter["1"]

		if zeros > ones {
			gamma_string.WriteString("0")
			epsilon_string.WriteString("1")
		} else {
			gamma_string.WriteString("1")
			epsilon_string.WriteString("0")
		}
	}

	gamma_rate := newBitstring(gamma_string.String())
	epsilon_rate := newBitstring(epsilon_string.String())
	fmt.Println("Part 1:", gamma_rate.value*epsilon_rate.value)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	part1()
	part2()
}
