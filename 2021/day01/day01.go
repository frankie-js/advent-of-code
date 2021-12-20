package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Println(strconv.Itoa(SolveDay01Part1(input)))
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
