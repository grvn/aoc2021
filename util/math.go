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
