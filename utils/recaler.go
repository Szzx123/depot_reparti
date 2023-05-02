package utils

// Recaler calcule l'horloge entière
func Recaler(x, y int) int {
	if x < y {
		return y + 1
	}
	return x + 1
}
