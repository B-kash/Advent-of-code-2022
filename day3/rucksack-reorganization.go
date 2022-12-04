package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Group struct {
	compartments      []string
	badgeWithPriority map[rune]int
}

func main() {
	rucksacks := getInputFromFile()
	totalPriority := 0
	fmt.Println("number of rucksacks:", len(rucksacks))
	for _, rucksack := range rucksacks {
		valuePrioMap := rucksack.badgeWithPriority
		for _, priotity := range valuePrioMap {
			totalPriority += priotity
		}

	}
	fmt.Println("My Total Priority", totalPriority)

}

func getInputFromFile() []Group {
	f, err := os.Open("day3/rucksack-reorganization.txt")
	var rucksacks []Group
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var compartments []string
		for i := 0; i < 3; i++ {
			line := scanner.Text()
			compartments = append(compartments, line)
			if i < 2 {
				scanner.Scan()
			}
		}
		rucksacks = append(rucksacks, Group{
			compartments:      compartments,
			badgeWithPriority: findCommonItem(compartments),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rucksacks

}

func findPriorityForItem(commonValue string) int {

	priorityMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	return priorityMap[commonValue]
}

func containsChar(s string, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findCommonItem(compartments []string) map[rune]int {
	compartment := compartments[0]
	commonItem := make(map[rune]int)
	for _, currentItem := range compartment {
		if containsChar(compartments[1], currentItem) && containsChar(compartments[2], currentItem) {
			commonItem[currentItem] = findPriorityForItem(string(currentItem))
		}
	}
	return commonItem
}
