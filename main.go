package main

import (
	"os"

	"github.com/grvn/aoc2021/cmd"
	log "github.com/sirupsen/logrus"
)

//go:generate go run ./gen
func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		log.WithFields((log.Fields{
			"Error": err,
		})).Error("Something crashed")
		os.Exit(1)
	}
}
