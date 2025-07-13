package main

func createMatrix(rows, cols int) [][]int {
	sliceOfSlices := make([][]int, rows)
	for i := range rows {
		sliceOfSlices[i] = make([]int, cols)
		for k := range cols {
			sliceOfSlices[i][k] = i * k
		}
	}
	return sliceOfSlices
}
