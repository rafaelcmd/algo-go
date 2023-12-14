package main

import "fmt"

func main() {
	var matrix = [][]int{
		{-7, -7},
		{100, 12},
		{-33, 17},
	}

	//Optimal Space & Time Complexity
	//O(w * h) time | O(w * h) space - where w is the width of the matrix and h is the height
	TransposeMatrix(matrix)
}

func TransposeMatrix(matrix [][]int) {
	numRows := len(matrix[0])
	numCols := len(matrix)
	var result [][]int

	for i := 0; i < numRows; i++ {
		row := make([]int, numCols)

		for j := 0; j < numCols; j++ {
			row[j] = matrix[j][i]
		}

		result = append(result, row)

		fmt.Println(result[i])
	}
}
