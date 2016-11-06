package matrix

// Inverse returns a inverse matrix of m.
// If there's no inverse matrix of m, nil is returned.
func Inverse(m [][]float64) [][]float64 {
	n := len(m)

	if n != len(m[0]) {
		return nil
	}

	mi := make([][]float64, n)

	for i := 0; i < n; i++ {
		mi[i] = make([]float64, n)
		mi[i][i] = 1.0
	}

	var mul float64

	for i := 0; i < n; i++ {
		if m[i][i] == 0.0 {
			return nil
		}

		mul = 1 / m[i][i]

		for j := 0; j < n; j++ {
			m[i][j] *= mul
			mi[i][j] *= mul
		}

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			mul = m[j][i]

			for k := 0; k < n; k++ {
				m[j][k] -= m[i][k] * mul
				mi[j][k] -= mi[i][k] * mul
			}
		}
	}

	return mi
}
