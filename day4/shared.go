package day4

import (
	"strings"
	"sync"

	"github.com/grvn/aoc2021/util"
)

type coordinat struct {
	x, y int
}

type board struct {
	numbers map[int]coordinat
	marked  map[int]map[int]bool
	size    int
}

type result struct {
	win   bool
	moves int
	value int
}

func createBoard(size int) board {
	markers := make(map[int]map[int]bool)
	for i := 0; i < size; i++ {
		markers[i] = map[int]bool{}
	}
	return board{
		numbers: map[int]coordinat{},
		marked:  markers,
		size:    size,
	}
}

func fillBoards(input []string) (boards []board) {
	y := 0
	length := len(strings.Fields(input[0]))
	board := createBoard(length)
	for _, line := range input {
		if line == "" {
			boards = append(boards, board)
			board = createBoard(length)
			y = 0
			continue
		}
		for x, value := range strings.Fields(line) {
			board.numbers[util.AtoIorEXIT(value)] = coordinat{x, y}
		}
		y++
	}
	boards = append(boards, board)
	return boards
}

func (board *board) tryToComplete(wg *sync.WaitGroup, drag []int, output chan result) {
	defer wg.Done()
	for i, number := range drag {
		if xy, ok := board.numbers[number]; ok {
			board.marked[xy.x][xy.y] = true
			if board.didIWin() {

				output <- result{
					win:   true,
					moves: i,
					value: number * board.calcUnmarked(),
				}
				return
			}
		}
	}
	output <- result{
		win: false,
	}
}

func (board *board) calcUnmarked() (output int) {
	for value, coord := range board.numbers {
		if !board.marked[coord.x][coord.y] {
			output += value
		}
	}
	return output
}

func (board *board) didIWin() bool {
noWinRow:
	for y := 0; y < board.size; y++ {
		for x := 0; x < board.size; x++ {
			if !board.marked[x][y] {
				continue noWinRow
			}
		}
		return true
	}

noWinCol:
	for x := 0; x < board.size; x++ {
		for y := 0; y < board.size; y++ {
			if !board.marked[x][y] {
				continue noWinCol
			}
		}
		return true
	}
	return false
}
