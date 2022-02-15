package day12

import (
	"testing"

	"github.com/grvn/aoc2021/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestSmall(t *testing.T) {
	viper.Set("input", "testS.txt")
	input := util.FromFile()
	result := execute1(input)
	require.Equal(t, 10, result)
}

func TestMedium(t *testing.T) {
	viper.Set("input", "testM.txt")
	input := util.FromFile()
	result := execute1(input)
	require.Equal(t, 19, result)
}

func TestLarge(t *testing.T) {
	viper.Set("input", "testL.txt")
	input := util.FromFile()
	result := execute1(input)
	require.Equal(t, 226, result)
}
