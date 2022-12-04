package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	totalCalorie int
}

func main() {
	elfs := getInputFromFile()
	highestCalorie := findTopThreeCalories(elfs)
	fmt.Println("Highest calorie: =", highestCalorie)
}

func findTopThreeCalories(elfs []Elf) int {
	sort.SliceStable(elfs, func(i, j int) bool {
		return elfs[i].totalCalorie > elfs[j].totalCalorie
	})
	return elfs[0].totalCalorie + elfs[1].totalCalorie + elfs[2].totalCalorie
}

func getInputFromFile() []Elf {
	f, err := os.Open("calorie-counting.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var allElfs []Elf
	scanner := bufio.NewScanner(f)
	var totalCalorie int = 0
	for scanner.Scan() {
		input := scanner.Text()
		calorie, err := strconv.Atoi(input)
		if err != nil {
			allElfs = append(allElfs, Elf{
				totalCalorie: totalCalorie,
			})
			totalCalorie = 0
			continue
		}
		totalCalorie += calorie

	}
	allElfs = append(allElfs, Elf{
		totalCalorie: totalCalorie,
	})
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allElfs
}
