package line

func generateEmptyVotes(longestDistance int, degrees int) [][]uint32 {
	result := make([][]uint32, longestDistance, longestDistance)

	for i := 0; i < len(result); i++ {
		result[i] = make([]uint32, degrees, degrees)
	}

	return result

}
