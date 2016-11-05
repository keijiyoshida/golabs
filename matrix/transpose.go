package matrix

// Transpose returns a transposed matrix of m.
func Transpose(m [][]float64) [][]float64 {
	mt := make([][]float64, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		mt[i] = make([]float64, len(m))
		for j := 0; j < len(m); j++ {
			mt[i][j] = m[j][i]
		}
	}

	return mt
}
