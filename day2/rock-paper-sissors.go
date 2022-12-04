package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type WinStatus int64

const (
	Win WinStatus = iota
	Loss
	Draw
)

type Move int64

const (
	ROCK Move = iota
	PAPER
	SISSOR
)

type Player struct {
	move       Move
	gameStatus WinStatus
	totalScore int
}

type Game struct {
	me  Player
	elf Player
}

func main() {
	games := getInputFromFile()
	totalScore := 0
	for _, game := range games {
		totalScore += game.me.totalScore
	}
	fmt.Println("My Total Score", totalScore)
}

func getMove(move string) Move {
	if contains(getRock(), move) {
		return ROCK
	}
	if contains(getPaper(), move) {
		return PAPER
	}
	return SISSOR
}

func getInputFromFile() []Game {
	f, err := os.Open("day2/rock-paper-sissors.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var allGames []Game
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		splits := strings.Split(input, " ")
		elfMove := getMove(splits[0])
		myWinStatus := getWinStatusForPlayer(splits[1])
		elfWinStatus := getOppositeWinStatus(myWinStatus)
		myMove := findMove(elfMove, myWinStatus)
		allGames = append(allGames, Game{
			me: Player{
				move:       myMove,
				gameStatus: myWinStatus,
				totalScore: calculateTotalScore(myMove, myWinStatus),
			},
			elf: Player{
				move:       elfMove,
				gameStatus: elfWinStatus,
				totalScore: calculateTotalScore(elfMove, elfWinStatus),
			},
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allGames
}

func getOppositeWinStatus(myWinStatus WinStatus) WinStatus {
	if myWinStatus == Win {
		return Loss
	} else if myWinStatus == Loss {
		return Win
	} else {
		return Draw
	}
}

func findMove(opponentMove Move, myWinStatus WinStatus) Move {
	switch opponentMove {
	case ROCK:
		if myWinStatus == Win {
			return PAPER
		} else if myWinStatus == Draw {
			return ROCK
		} else if myWinStatus == Loss {
			return SISSOR
		}
		break
	case PAPER:
		if myWinStatus == Win {
			return SISSOR
		} else if myWinStatus == Draw {
			return PAPER
		} else if myWinStatus == Loss {
			return ROCK
		}
		break
	case SISSOR:
		if myWinStatus == Win {
			return ROCK
		} else if myWinStatus == Loss {
			return PAPER
		} else if myWinStatus == Draw {
			return SISSOR
		}
	default:
		return ROCK
	}
	return ROCK
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getRock() []string {
	return []string{"A"}
}
func getPaper() []string {
	return []string{"B"}
}
func getSissor() []string {
	return []string{"C"}
}

func getWinStatusForPlayer(winStatus string) WinStatus {
	if winStatus == "X" {
		return Loss
	}
	if winStatus == "Z" {
		return Win
	}
	return Draw
}

func checkWinCondition(yourMove Move, opponentMove Move) WinStatus {
	gameStatus := Draw
	if yourMove == ROCK {
		if opponentMove == SISSOR {
			return Win
		} else if opponentMove == PAPER {
			return Loss
		}
	} else if yourMove == PAPER {
		if opponentMove == SISSOR {
			return Loss
		} else if opponentMove == ROCK {
			return Win
		}
	} else {
		if opponentMove == PAPER {
			return Win
		} else if opponentMove == ROCK {
			return Loss
		}
	}
	return gameStatus
}

func getScoreByWinCondition(winStatus WinStatus) int {
	score := 0
	if winStatus == Win {
		return 6
	}

	if winStatus == Draw {
		return 3
	}

	return score
}

func getScoreByMove(move Move) int {
	score := 1
	if move == PAPER {
		return 2
	}
	if move == SISSOR {
		return 3
	}
	return score
}

func calculateTotalScore(yourMove Move, winStatus WinStatus) int {
	return getScoreByMove(yourMove) + getScoreByWinCondition(winStatus)
}
