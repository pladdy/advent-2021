package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("day-1-puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	increases := 0
	lastMeasurement := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = line[:len(line)-1] // remove new line char
		measurement, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// don't count the first measurement as an increase
		if measurement > lastMeasurement && lastMeasurement > 0 {
			increases++
		}
		lastMeasurement = measurement
	}
	fmt.Println("Part 1:", increases)
}

func part2() {
	file, err := os.Open("day-1-puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create windows
	reader := bufio.NewReader(file)
	windows := [][]int{}
	windowA := []int{}
	windowB := []int{}
	windowC := []int{}

	i := 0
	for {
		if len(windowA) == 3 {
			windows = append(windows, windowA)
			windowA = []int{}
		}

		if len(windowB) == 3 {
			windows = append(windows, windowB)
			windowB = []int{}
		}

		if len(windowC) == 3 {
			windows = append(windows, windowC)
			windowC = []int{}
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = line[:len(line)-1] // remove new line char
		measurement, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// windows can only have 3 items at a time
		if len(windowC) < len(windowB) || len(windowB) == 0 {
			windowC = append(windowC, measurement)
		}
		if len(windowB) < len(windowA) || len(windowA) == 0 {
			windowB = append(windowB, measurement)
		}
		if len(windowA) < 3 {
			windowA = append(windowA, measurement)
		}

		i++
	}

	// sum windows
	lastSum := 0
	increases := 0

	for _, window := range windows {
		sum := 0
		for _, measure := range window {
			sum += measure
		}
		if sum > lastSum && lastSum > 0 {
			increases++
		}
		lastSum = sum
	}

	// 466 too low
	fmt.Println("Part 2:", increases)
}

func main() {
	part1()
	part2()
}
