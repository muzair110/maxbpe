package utils

import (
	"github.com/muzair110/maxbpe/schema"
)

func MintTokens(tokens []int32,
	mostRepeatingPair schema.Tuple,
	newTokenID int32,
) []int32 {
	/*
		This function is responsible for minting new tokens based on the most repeating byte sequence
			Input arguments:
				-- tokens: A slice of raw byte sequence
				-- mostRepeatingPair : The most repeating byte sequence in the raw bytes
				-- newTokenID : This is the token ID that is to be minted for the most repeating pair

			Output arguments:
				-- A slice of byte sequence with newly minted tokens
	*/
	idx := 0
	var newTokens []int32

	for idx < len(tokens) {
		// In case we encounter the most repeating pair, we are minting a new token
		if idx < len(tokens)-1 && tokens[idx] == mostRepeatingPair.X && tokens[idx+1] == mostRepeatingPair.Y {
			newTokens = append(newTokens, newTokenID)
			idx += 2
		} else {
			newTokens = append(newTokens, tokens[idx])
			idx += 1
		}

	}
	return newTokens
}
