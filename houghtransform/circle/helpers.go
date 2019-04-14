package circle

func generateEmptyVotes(maxRadius, maxA, maxB int) [][][]uint32 {
	result := make([][][]uint32, maxA, maxA)

	for i := 0; i < len(result); i++ {
		result[i] = make([][]uint32, maxB, maxB)

		for j := 0; j < len(result[i]); j++ {
			result[i][j] = make([]uint32, maxRadius, maxRadius)
		}
	}

	return result

}
