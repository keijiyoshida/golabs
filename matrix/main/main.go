package main

import (
	"fmt"

	"github.com/keijiyoshida/golabs/matrix"
)

func main() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	mt := matrix.Transpose(m)
	fmt.Println(m)
	fmt.Println(mt)
}
