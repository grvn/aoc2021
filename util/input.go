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

func (i *Input) GetIntSlice() (input []int) {
	for line := range i.lines {
		intline, err := strconv.Atoi(line)
		if err != nil {
			log.WithFields((log.Fields{
				"Error": err,
			})).Error("line can not be converted to int")
			os.Exit(1)
		}
		input = append(input, intline)
	}
	return input
}
