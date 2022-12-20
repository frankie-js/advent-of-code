package main

import (
	"log"
	"os"
	"strings"
)

// Rock 		= A, 1
// Paper 		= B, 2
// Scissors = C, 3
// Lose, Draw, Win = 0 X, 3 Y, 6 Z

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	splitStr := strings.Split(string(dat), "\n")

	var actions [][]string

	for _, line := range splitStr {
		if line == "" {
			continue
		}
		split := strings.Split(string(line), " ")
		actions = append(actions, split)
	}

	total := 0

	for _, round := range actions {
		if round[1] == "X" {
			if round[0] == "A" {
				total += 3
			} else if round[0] == "B" {
				total += 1
			} else if round[0] == "C" {
				total += 2
			}
		} else if round[1] == "Y" {
			if round[0] == "A" {
				total += 3 + 1
			} else if round[0] == "B" {
				total += 3 + 2
			} else if round[0] == "C" {
				total += 3 + 3
			}
		} else if round[1] == "Z" {
			if round[0] == "A" {
				total += 6 + 2
			} else if round[0] == "B" {
				total += 6 + 3
			} else if round[0] == "C" {
				total += 6 + 1
			}
		}
	}

	log.Println(total)
}
