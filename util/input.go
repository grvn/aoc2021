package util

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Input struct {
	// use bufio to read files
	// https://gobyexample.com/reading-files
	scanner *bufio.Scanner
	lines   chan string
}

func FromFile() *Input {
	path := viper.GetString("input")
	if path == "" {
		ex, err := os.Executable()
		if err != nil {
			log.WithFields((log.Fields{
				"Error": err,
			})).Error("Can not get input, use -i")
			os.Exit(1)
		}
		path = filepath.Join(filepath.Dir(ex), "input.txt")
	}
	file, err := os.Open(path)
	if err != nil {
		log.WithFields((log.Fields{
			"Error": err,
		})).Error("Can not open input file")
		os.Exit(1)
	}
	input := &Input{
		scanner: bufio.NewScanner(file),
		lines:   make(chan string),
	}
	// https://stackoverflow.com/questions/27217428/reading-a-file-concurrently
	go func() {
		for input.scanner.Scan() {
			input.lines <- input.scanner.Text()
		}
		close(input.lines)
	}()
	return input
}

func (input *Input) GetIntSlice() (output []int) {
	for line := range input.GetLines() {
		intline, err := strconv.Atoi(line)
		if err != nil {
			log.WithFields((log.Fields{
				"Error": err,
			})).Error("line can not be converted to int")
			os.Exit(1)
		}
		output = append(output, intline)
	}
	return output
}

func (input *Input) GetStringSlice() (output []string) {
	for line := range input.GetLines() {
		output = append(output, line)
	}
	return output
}

func (input *Input) GetIntMatrix() (output [][]int) {
	lines := input.GetStringSlice()
	cols := len(lines[0])
	rows := len(lines)

	output = make([][]int, rows)
	for i := range output {
		output[i] = make([]int, cols)
	}
	for i, line := range lines {
		for j, value := range line {
			// subract the UTF-8 representation of 0 to get the "actual" number
			output[i][j] = int(value - '0')
		}
	}
	return output
}

func (input *Input) GetLines() <-chan string {
	return input.lines
}
