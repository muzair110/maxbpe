package utils

import "github.com/muzair110/maxbpe/schema"

func GetMostFrequentPair(tokens []int32) (schema.Tuple, int32) {

	/*
		This function is responsible for finding the most frequent pair from the byte sequence
			Input arguments:
				-- tokens: A slice of raw byte sequence

			Output arguments:
				-- A tuple that is the most frequence pair
				-- Number of times the pair appeared in the raw tokens
	*/

	counts := make(map[schema.Tuple]int32)

	idx := 0
	for idx < len(tokens)-1 {
		pair := schema.Tuple{X: tokens[idx], Y: tokens[idx+1]}

		// Increments the total count
		counts[pair] = counts[pair] + 1
		idx += 1
	}

	// Finding the most frequent pair
	var maxCount int32 = 0
	mostFrequentPair := schema.Tuple{X: 0, Y: 0}

	for key, val := range counts {

		if val > maxCount {
			maxCount = val
			mostFrequentPair = key
		}
	}
	return mostFrequentPair, maxCount
}
