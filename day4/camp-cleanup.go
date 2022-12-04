package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"Advent/utils"
)

type ElfGroup struct {
	workList       [][]int
	hasFullOverlap bool
}

func main() {
	elfGroup := getInputFromFile()
	freeElfs := 0
	for _, elf := range elfGroup {
		if elf.hasFullOverlap {
			freeElfs++
		}
	}
	fmt.Println("Number of overlaps:", freeElfs)
}

func getInputFromFile() []ElfGroup {
	f, err := os.Open("day4/camp-cleanup.txt")

	if err != nil {
		log.Fatal(err)
	}
	var elfGroup []ElfGroup
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ",")
		var workList [][]int
		for _, workLoad := range splits {
			workList = append(workList, getWorkListFromRange(workLoad))
		}
		elfGroup = append(elfGroup, ElfGroup{
			workList:       workList,
			hasFullOverlap: checkWorkOverlap(workList),
		})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return elfGroup
}

func checkWorkOverlap(workList [][]int) bool {
	overlapExists := false

	for i := 0; i < len(workList[0]) && !overlapExists; i++ {
		if utils.ContainsInIntArray(workList[1], workList[0][i]) {
			overlapExists = true
		}
	}

	return overlapExists

}

func getWorkListFromRange(workRange string) []int {
	splits := strings.Split(workRange, "-")
	startRange, _ := strconv.Atoi(splits[0])
	endRange, _ := strconv.Atoi(splits[1])
	var workList []int
	for i := startRange; i <= endRange; i++ {
		workList = append(workList, i)
	}
	return workList
}
