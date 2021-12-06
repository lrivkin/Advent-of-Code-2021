package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Spot  struct {
	num string
	marked bool
}

type Board struct {
	spots []Spot
	won bool
}

type Game struct {
	winningPositions [][]int
	boards []Board
	moves []string
	numWinners int
}

func initializeGame(game *Game) {
	var (
		row []int
		col []int
		winningPatterns [][]int
	)

	// Various combos of winning row positions
	for j:=0; j<5; j++ {
		row = []int{0, 0, 0, 0, 0}
		col = []int{0, 0, 0, 0, 0}
		for i:=0; i<5; i++ {
			row[i] = i + j*5
			col[i] = i*5 + j
		}
		winningPatterns = append(winningPatterns, row)
		winningPatterns = append(winningPatterns, col)
	}
	//diag1 := []int{0, 0, 0, 0, 0}
	//diag2 := []int{0, 0, 0, 0, 0}
	//for i:= 0; i<5; i++ {
	//	diag1[i] = i + i*5
	//	diag2[i] = 4-i + i*5
	//}
	//// diagonals. I'm sure there's an elegant way to get these but :shrug:
	//winningPatterns = append(winningPatterns, diag1)
	//winningPatterns = append(winningPatterns, diag2)
	game.winningPositions = winningPatterns
	game.numWinners = 0
}


func hasWon(game *Game, boardNum int) bool {
	board := game.boards[boardNum]
	for _, possibleCombo := range game.winningPositions {
		allMatch := true
		for _, pos := range possibleCombo {
			if board.spots[pos].marked == false {
				allMatch = false
				break
			}
		}
		if allMatch {
			return true
		}
	}
	return false
}

func checkGameResults(game *Game) int {
	for i := range game.boards {
		if !game.boards[i].won {
			if hasWon(game, i) {
				game.boards[i].won = true
				game.numWinners++
				if game.numWinners == len(game.boards) {
					return i
				}
			}
		}
		//fmt.Printf("Board %v\n", i)
		//printBoard(game.boards[i])
	}
	//fmt.Println()
	//fmt.Printf("%d boards have won so far\n", game.numWinners)
	return -1
}

func playMove(move string, game *Game) bool {
	//fmt.Printf("Move: %s \n", move)
	for _, board := range game.boards {
		for i, spot := range board.spots {
			if spot.num == move {
				board.spots[i].marked = true
				//board.score = board.score+1
				break
			}
		}
		//fmt.Println(board)
	}
	return false
}

func finalScore(board Board, move string) int {
	//	Start by finding the sum of all unmarked numbers on that board;
	//	in this case, the sum is 188.
	//	Then, multiply that sum by the number that was just called
	//		when the board won, 24, to get the final score, 188 * 24 = 4512.
	unmarkedSum := 0
	for _, spot := range board.spots {
		if !spot.marked {
			spotVal, _ := strconv.Atoi(spot.num)
			unmarkedSum += spotVal
		}
	}
	moveVal, _ := strconv.Atoi(move)
	return unmarkedSum*moveVal
}

func printBoard(board Board) {
	row := make([]string, 5)
	for i, s := range board.spots {
		val := s.num
		if n, _ := strconv.Atoi(val); n < 10 {
			val = " " + val
		}
		if s.marked {
			row[i%5] = val + "*"
		} else {
			row[i%5] = val + " "
		}
		if i%5 == 4 {
			fmt.Println(row)
		}
	}
}

func readBoard(path string, game *Game) {
	contents, _ := ioutil.ReadFile(path)
	parts := strings.Split(string(contents), "\n\n")
	moves := strings.Split(parts[0], ",")
	game.moves = moves
	var boards []Board
	var pieces []Spot
	for _, part := range parts[1:] {
		pieces = []Spot{}
		splitUp := strings.Split(part, "\n")
		for _, line := range splitUp {
			s := strings.Split(line, " ")
			for _, x := range s {
				x = strings.TrimSpace(x)
				if x != "" {
					pieces = append(pieces,Spot{x, false})
				}
			}
		}
		boards = append(boards, Board{pieces, false})
	}
	game.boards = boards

}



func main() {
	fmt.Println("Advent of Code Day 4")
	var game Game
	initializeGame(&game)
	//fmt.Println(game.winningPositions)
	readBoard("inputs/day4.txt", &game)
	for _, move := range game.moves {
		playMove(move, &game)
		//fmt.Printf("Move %s \n", move)
		loser := checkGameResults(&game)
		if loser != -1 {
			printBoard(game.boards[loser])
			fmt.Println(finalScore(game.boards[loser], move))
			break
		}
	}
}

