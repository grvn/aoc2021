package day8

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

// Testing using buffered goroutines
// and workerpools

// Segments
// 0 -> 6
// 1 -> 2
// 2 -> 5
// 3 -> 5
// 4 -> 4
// 5 -> 5
// 6 -> 6
// 7 -> 3
// 8 -> 7
// 9 -> 6

// Number of segments
// 2 -> 1
// 3 -> 7
// 4 -> 4
// 5 -> 2,3,5
// 6 -> 0,6,9
// 7 -> 8

const (
	WORKERS = 8
)

var jobs = make(chan Job, WORKERS)
var results = make(chan int, WORKERS)

type Job struct {
	leftSide  []string
	rightside []string
}

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 8, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day8, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	sum := make(chan int)
	go createJobs(input)
	go result(sum)
	createWorkers()
	return <-sum
}

func result(sum chan int) {
	add := 0
	for result := range results {
		add += result
	}
	sum <- add
}

func createWorkers() {
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := solveLine(job)
		results <- output
	}
	wg.Done()
}

func createJobs(input *util.Input) {
	i := 0
	for line := range input.GetLines() {
		splitLine := strings.Split(line, " | ")
		leftSideParts := strings.Fields(splitLine[0])
		rightSideParts := strings.Fields(splitLine[1])
		job := Job{
			leftSide:  leftSideParts,
			rightside: rightSideParts,
		}
		jobs <- job
		i++
	}
	close(jobs)
}

func solveLine(job Job) int {
	code := map[int]string{}
	solution := 0
	five := []string{}
	six := []string{}
	for _, section := range job.leftSide {
		switch len(section) {
		case ONE:
			code[1] = util.SortString(section)
		case FOUR:
			code[4] = util.SortString(section)
		case SEVEN:
			code[7] = util.SortString(section)
		case EIGHT:
			code[8] = util.SortString(section)
		case 5:
			five = append(five, util.SortString(section))
		case 6:
			six = append(six, util.SortString(section))
		}
	}
	// Get number 9, which contains all the same as 4
	// number 0 and 6 does not contain 4
	for i, section := range six {
		if containsRunes([]rune(code[4]), section) {
			code[9] = section
			six = removeFromSlice(six, i)
			break
		}
	}
	// Get number 0 and 6
	// 0 contains all of 1 which 6 does not
	for i, section := range six {
		if containsRunes([]rune(code[1]), section) {
			code[0] = section
			six = removeFromSlice(six, i)
			break
		}
	}
	code[6] = six[0]
	// Get number 3, which contains all the same as 1
	// which 2 and 5 does not
	for i, section := range five {
		if containsRunes([]rune(code[1]), section) {
			code[3] = section
			five = removeFromSlice(five, i)
			break
		}
	}
	// Get number 5, which contains all the same as 6
	// which 2 does not
	// but reverse it because 6 contains one more than 5
	for i, section := range five {
		if containsRunes([]rune(section), code[6]) {
			code[5] = section
			five = removeFromSlice(five, i)
			break
		}
	}
	code[2] = five[0]
	for _, display := range job.rightside {
		// shift numbers right when using decimal
		solution *= 10
		for k, v := range code {
			if v == util.SortString(display) {
				solution += k
			}
		}
	}
	return solution
}

func containsRunes(runes []rune, s string) bool {
	for _, rune := range runes {
		if !strings.ContainsRune(s, rune) {
			return false
		}
	}
	return true
}

func removeFromSlice(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1] // Copy last element to index i.
	slice[len(slice)-1] = ""           // Erase last element (write zero value).
	slice = slice[:len(slice)-1]
	return slice
}
