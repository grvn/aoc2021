package day10

import (
	"testing"

	"github.com/grvn/aoc2021/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	viper.Set("input", "test.txt")
	input := util.FromFile()
	result := execute1(input)
	require.Equal(t, 26397, result)
}
