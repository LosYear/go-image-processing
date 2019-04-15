package line

type localMaxima2D struct {
	x, y  int
	value uint32
}

func findLocalMaxima(array *[][]uint32, threshold uint32) []localMaxima2D {
	var maximas []localMaxima2D

	for i := 0; i < len(*array); i++ {
		for j := 0; j < len((*array)[i]); j++ {
			value := (*array)[i][j]
			flag := true

			if value <= threshold {
				continue
			}

			if (i-1 >= 0 && j-1 >= 0 && value < (*array)[i-1][j-1]) ||
				(i-1 >= 0 && value < (*array)[i-1][j]) ||
				(j-1 >= 0 && value < (*array)[i][j-1]) ||
				(i+1 < len(*array) && value < (*array)[i+1][j]) ||
				(j+1 < len((*array)[i]) && value < (*array)[i][j+1]) ||
				(i+1 < len(*array) && j+1 < len((*array)[i]) && value < (*array)[i+1][j+1]) ||
				(i-1 >= 0 && j+1 < len((*array)[i]) && value < (*array)[i-1][j+1]) ||
				(j-1 >= 0 && i+1 < len(*array) && value < (*array)[i+1][j-1]) {
				flag = false
			}

			if flag {
				maximas = append(maximas, localMaxima2D{x: i, y: j, value: value})
			}

		}
	}

	return maximas

}
