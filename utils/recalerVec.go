package utils

func RecalerVec(x, y []int) []int {
	res := make([]int, 3)
	res[1] = max(x[1], y[1])
	res[2] = max(x[2], y[2])
	res[3] = max(x[3], y[3])
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
