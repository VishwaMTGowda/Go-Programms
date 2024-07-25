package main

import (
	"fmt"
)

func main() {
	var matrix1 [100][100]int
	var matrix2 [100][100]int
	var sum [100][100]int
	var row, col int

	fmt.Println("Enter number of rows:")
	fmt.Scanln(&row)
	fmt.Println("Enter number of columns:")
	fmt.Scanln(&col)

	fmt.Println("\n========== Matrix 1 ==========")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for matrix1 %d%d: ", i+1, j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}

	fmt.Println("\n========== Matrix 2 ==========")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for matrix2 %d%d: ", i+1, j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	// Compute the sum of matrices
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	fmt.Println("\n========== Sum of Matrices ==========")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%d\t", sum[i][j])
		}
		fmt.Println()
	}
}
