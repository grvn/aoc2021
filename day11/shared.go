package day11

import "fmt"

func getNeighbours(points map[point]int, p point) (neighbours []point) {
	var neighbourpoints = []struct {
		x, y int
	}{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y + 1},
	}

	for _, point := range neighbourpoints {
		if _, ok := points[point]; ok {
			neighbours = append(neighbours, point)
		}
	}
	return neighbours
}

func printGrid(m map[point]int) {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			v, ok := m[point{x: x, y: y}]
			if ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
