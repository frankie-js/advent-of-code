package main

import (
	"log"
	"os"
	"sort"
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

	var elfTotals []int
	for _, elf := range elves {
		elfTotal := 0
		for _, v := range elf {
			elfTotal += v
		}
		elfTotals = append(elfTotals, elfTotal)
	}

	sort.Ints(elfTotals)

	largestElves := [3]int{elfTotals[len(elfTotals)-1], elfTotals[len(elfTotals)-2], elfTotals[len(elfTotals)-3]}
	largestElvesTotal := 0
	for _, v := range largestElves {
		largestElvesTotal += v
	}

	log.Println("Largest Elves Total", largestElvesTotal)
}
