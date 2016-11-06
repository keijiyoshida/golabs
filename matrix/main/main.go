package main

import (
	"fmt"

	"github.com/keijiyoshida/golabs/matrix"
)

func main() {
	m1 := [][]float64{
		{1.0, 2.0, 0.0, -1.0},
		{-1.0, 1.0, 2.0, 0.0},
		{2.0, 0.0, 1.0, 1.0},
		{1.0, -2.0, -1.0, 1.0},
	}

	fmt.Println(m1)
	fmt.Println(matrix.Transpose(m1))

	m2 := [][]float64{
		{1.0, 1.0, -1.0},
		{-2.0, 0.0, 1.0},
		{0.0, 2.0, 1.0},
	}

	fmt.Println(matrix.Inverse(m1))
	fmt.Println(matrix.Inverse(m2))
}
