package day9

import (
	"testing"

	"github.com/grvn/aoc2021/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func Test2(t *testing.T) {
	viper.Set("input", "test.txt")
	input := util.FromFile()
	result := execute2(input)
	require.Equal(t, 1134, result)
}
