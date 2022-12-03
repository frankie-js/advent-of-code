package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := readLines("input.txt")

	if err != nil {
		panic(err)
	}

	var input = []int{}

	for _, i := range dat {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		input = append(input, j)
	}

	fmt.Println("Part 1:", strconv.Itoa(SolveDay01Part1(input)))
	// fmt.Println("Part 2:", strconv.Itoa(SolveDay01Part2(input)))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SolveDay01Part1(input []int) int {
	count := 0

	for i := range input {
		if i == 0 {
			continue
		}

		if input[i-1] < input[i] {
			count++
		}
	}

	return count
}

// func SolveDay01Part2(input []int) int {
// 	count := 0
// 	previousSum := 0

// 	for i := range input {
// 		if i >= 0 && i <= 2 {
// 			continue
// 		}

// 		sum := input[i] + input[i-1] + input[i-2]

// 		if previousSum < sum {
// 			count++
// 			previousSum = sum
// 		}
// 	}

// 	return count
// }

func InputToIntSlice(input string) ([]int, error) {
	var i []int
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		i = append(i, lineInt)
	}
	return i, nil
}

func SolveDay01Part2(input string) int {
	inputInt, _ := InputToIntSlice(input)
	avg := make([]int, len(inputInt)-2)
	for i := range avg {
		avg[i] += inputInt[i]
		avg[i] += inputInt[i+1]
		avg[i] += inputInt[i+2]
	}
	return calculateIncrements(avg)
}

func calculateIncrements(input []int) int {
	var increments int
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] > input[i-1] {
			increments++
		}
	}
	return increments
}
