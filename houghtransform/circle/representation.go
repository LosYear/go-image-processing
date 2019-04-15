package circle

func codeCenter(a, b, maxX, maxY int) (int, int) {
	return a + maxX, b + maxY
}

func decodeCenter(a, b, maxX, maxY int) (int, int) {
	return a - maxX, b - maxY
}
