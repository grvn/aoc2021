package day5

type board map[int]map[int]int

func (board board) leastDanger() (output int) {
	for _, col := range board {
		for _, val := range col {
			if val > 1 {
				output++
			}
		}
	}
	return output
}
