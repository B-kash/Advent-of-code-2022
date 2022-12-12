package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack = []string

type Shipment struct {
	pillers []Stack
	moves   []Move
}

type Move struct {
	times int
	from  int
	to    int
}

func unShift(array []string, data string) []string {
	array = append([]string{data}, array...)
	return array
}

func push(array []string, data []string) []string {
	array = append(array, data...)
	return array
}

func (s *Shipment) moveShipment() {
	for _, move := range s.moves {
		fromPiller := s.pillers[move.from]
		toPiller := s.pillers[move.to]
		lowerBound := len(fromPiller) - move.times

		if (len(fromPiller)-1 < 0) || (len(fromPiller)-move.times < 0) {
			lowerBound = 0
		}
		top := fromPiller[lowerBound:]
		s.pillers[move.from] = fromPiller[0:lowerBound]
		toPiller = append(toPiller, top...)
		s.pillers[move.to] = toPiller
	}
}

func main() {
	shipment := getInputFromFile()
	// fmt.Println(shipment.pillers)
	// fmt.Println(shipment.moves)
	shipment.moveShipment()
	topCrates := ""
	for _, piller := range shipment.pillers {
		if len(piller) == 0 {
			continue
		}
		top := piller[len(piller)-1]
		piller = piller[:len(piller)-1]
		topCrates += top
	}
	fmt.Println("Top Crates:", topCrates)
}

func getInputFromFile() Shipment {
	f, err := os.Open("day5/supply-stacks.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	shipmentCompleted := false
	var shipment Shipment
	hasShipmentSize := false
	moves := make([]Move, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			shipmentCompleted = true
		}
		if !hasShipmentSize {
			shipment.pillers = make([]Stack, (len(line)/4)+1)
			hasShipmentSize = true
		}
		for i := 0; i < len(line) && !shipmentCompleted; i += 4 {
			endIndex := i + 3

			if endIndex >= len(line) {
				endIndex = len(line) - 1
			}
			crate := strings.Trim(line[i:endIndex], " []")
			if len(crate) == 0 {
				continue
			}
			pillerNumber, err := strconv.Atoi(crate)

			if err == nil && pillerNumber != 0 {
				break
			}

			index := i / 4
			shipment.pillers[index] = unShift(shipment.pillers[index], crate)
		}
		if shipmentCompleted {
			if len(line) == 0 {
				continue
			}
			move := gatherMoves(line)
			moves = append(moves, move)
		}

	}
	shipment.moves = moves
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return shipment
}

func gatherMoves(line string) Move {

	splits := strings.Split(line, " ")
	times, _ := strconv.Atoi(splits[1])
	from, _ := strconv.Atoi(splits[3])
	to, _ := strconv.Atoi(splits[5])
	return Move{
		times: times,
		from:  from - 1,
		to:    to - 1,
	}
}
