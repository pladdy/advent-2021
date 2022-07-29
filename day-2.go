package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("day-2-puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	depth := 0
	horizon_position := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1] // chomp
		values := strings.Fields(line)

		direction, distance_str := values[0], values[1]
		distance, err := strconv.Atoi(distance_str)
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizon_position += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	fmt.Println("Part 1: ", depth*horizon_position)
}

func part2() {
	file, err := os.Open("day-2-puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	aim := 0
	depth := 0
	horizon_position := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1] // chomp
		values := strings.Fields(line)

		direction, distance_str := values[0], values[1]
		distance, err := strconv.Atoi(distance_str)
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizon_position += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	fmt.Println("Part 2: ", depth*horizon_position)
}

func main() {
	part1()
	part2()
}
