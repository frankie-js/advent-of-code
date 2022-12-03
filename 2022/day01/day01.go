package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	splitStr := strings.Split(string(dat), "\n\n")

	var elves [][]int
	for _, elfStr := range splitStr {
		subSplitStr := strings.Split(elfStr, "\n")

		var intList []int
		for _, item := range subSplitStr {
			if item == "" {
				continue
			}
			k, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalln(err)
			}
			intList = append(intList, k)
		}

		elves = append(elves, intList)
	}

	var largestElfTotal int
	for _, elf := range elves {
		result := 0
		for _, v := range elf {
			result += v
		}
		if result > largestElfTotal {
			largestElfTotal = result
		}
	}

	log.Println("Largest Elf Total", largestElfTotal)
}
