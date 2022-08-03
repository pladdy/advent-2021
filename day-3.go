package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bitStringToInt(s string) int64 {
	value, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return value
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

	gamma_rate := bitStringToInt(gamma_string.String())
	epsilon_rate := bitStringToInt(epsilon_string.String())
	fmt.Println("Part 1:", gamma_rate*epsilon_rate)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	part1()
	part2()
}
