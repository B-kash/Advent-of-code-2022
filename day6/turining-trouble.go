package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type device struct {
	dataStream string
}

var lengthOfPacket = 4
var lenghthOfMessage = 14

func isNonRepeatingSequence(dataStream string) bool {
	isNonRepeatingSequence := true
	for _, d := range dataStream {
		count := strings.Count(dataStream, string(d))
		if count > 1 {
			isNonRepeatingSequence = false
			break
		}
	}
	return isNonRepeatingSequence
}

func (d *device) detectStartOfPacketMarker() int {
	startOfPacketMarker := 0
	for i := 0; i < len(d.dataStream)-lengthOfPacket; i++ {
		endOfChunk := i + lengthOfPacket
		if isNonRepeatingSequence(d.dataStream[i:endOfChunk]) {
			startOfPacketMarker = endOfChunk
			break
		}
	}
	return startOfPacketMarker
}
func (d *device) detectStartOfMessageMarker() int {
	startOfMessageMarker := 0
	for i := 0; i < len(d.dataStream)-lenghthOfMessage; i++ {
		endOfChunk := i + lenghthOfMessage
		if isNonRepeatingSequence(d.dataStream[i:endOfChunk]) {
			startOfMessageMarker = endOfChunk
			break
		}
	}
	return startOfMessageMarker
}

func main() {
	input := getInputFromFile()
	device := &device{
		dataStream: input,
	}
	startPacketMarker := device.detectStartOfPacketMarker()
	startMessageMarker := device.detectStartOfMessageMarker()

	fmt.Println(startPacketMarker, startMessageMarker)
}

func getInputFromFile() string {
	f, err := os.Open("day6/turning-trouble.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}
