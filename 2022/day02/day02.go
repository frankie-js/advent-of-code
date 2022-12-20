package main

import (
	"log"
	"os"
	"strings"
)

// Rock 		= A, X, 1
// Paper 		= B, Y, 2
// Scissors = C, Z, 3
// Win, Draw, Lose = 6, 3, 0

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
		roundTotal := 0
		if round[1] == "X" {
			roundTotal += 1
		} else if round[1] == "Y" {
			roundTotal += 2
		} else if round[1] == "Z" {
			roundTotal += 3
		}

		if round[0] == "A" {
			if round[1] == "X" {
				roundTotal += 3 // Draw
			} else if round[1] == "Y" {
				roundTotal += 6 // Win
			}
		} else if round[0] == "B" {
			if round[1] == "Y" {
				roundTotal += 3 // Draw
			} else if round[1] == "Z" {
				roundTotal += 6 // Win
			}
		} else if round[0] == "C" {
			if round[1] == "Z" {
				roundTotal += 3 // Draw
			} else if round[1] == "X" {
				roundTotal += 6 // Win
			}
		}

		total += roundTotal
	}

	log.Println(total)
}
