package utils

// RecalerVec calcule l'horloge vectorielle
func RecalerVec(x, y []int) []int {
	res := make([]int, 3)
	res[0] = max(x[0], y[0])
	res[1] = max(x[1], y[1])
	res[2] = max(x[2], y[2])
	return res
}

// max retourne l'entier le plus grand
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
