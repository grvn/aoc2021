package util

func Transpose(input [][]int) (output [][]int) {
	cols := len(input[0])
	rows := len(input)
	output = make([][]int, cols)
	for i := range output {
		output[i] = make([]int, rows)
	}
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			output[col][row] = input[row][col]
		}
	}
	return output
}

func RotateRight(array []int, steps int) []int {
	length := len(array)
	if length > 0 {
		if steps > length {
			steps = steps % length
		}
		array = append(array[length-steps:length], array[0:length-steps]...)
	}
	return array
}

func RotateLeft(array []int, steps int) []int {
	length := len(array)
	if length > 0 {
		if steps > length {
			steps = steps % length
		}
		array = append(array[steps:], array[0:steps]...)
	}
	return array
}
