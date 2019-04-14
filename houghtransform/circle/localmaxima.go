package circle

type localMaxima3D struct {
	x, y, z int
	value   uint32
}

func findLocalMaxima(array *[][][]uint32, threshold uint32) []localMaxima3D {
	var maximas []localMaxima3D

	maxI := len(*array)
	maxJ := len((*array)[0])
	maxK := len((*array)[0][0])
	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			for k := 0; k < maxK; k++ {
				value := (*array)[i][j][k]

				if value < threshold {
					continue
				}

				flag := true

				for iOffset := -1; flag && iOffset <= 1; iOffset++ {
					for jOffset := -1; flag && jOffset <= 1; jOffset++ {
						for kOffset := -1; flag && kOffset <= 1; kOffset++ {
							if iOffset == 0 && jOffset == 0 && kOffset == 0 {
								continue
							}

							iOffseted := i + iOffset
							jOffseted := j + jOffset
							kOffseted := k + kOffset

							if iOffseted >= 0 && jOffset >= 0 && kOffseted >= 0 &&
								iOffseted < maxI && kOffseted < maxK && jOffseted < maxJ &&
								value < (*array)[iOffseted][jOffseted][kOffseted] {
								flag = false

							}
						}
					}
				}

				if flag {
					maximas = append(maximas, localMaxima3D{x: i, y: j, z: k, value: value})
				}
			}
		}
	}
	return maximas
}
