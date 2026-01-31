package bpe

import (
	"fmt"
	"log"

	"github.com/muzair110/maxbpe/schema"
	"github.com/muzair110/maxbpe/utils"
)

type Tokenizer struct {
	schema.Assets
	ShowDebugLogs bool
}

func (t *Tokenizer) Train(vocabSize, startToken int32, text string) {
	if t.ShowDebugLogs {
		fmt.Println("========== TRAINING STARTED ==========")
		fmt.Printf("Vocab Size: %d, Start Token: %d, Text Length: %d chars\n", vocabSize, startToken, len(text))
	}

	if startToken > vocabSize {
		fmt.Println("ERROR: startToken > vocabSize, returning early")
		return
	}

	// This avoids PANIC
	if t.Merges == nil {
		if t.ShowDebugLogs {
			fmt.Println("Initializing Merges map...")
		}

		t.Merges = make(map[schema.Tuple]int32)
	}

	if t.Vocabulary == nil {
		if t.ShowDebugLogs {
			fmt.Println("Initializing Vocabulary map...")
		}

		t.Vocabulary = make(map[int32]string)
		var idx int32

		// Adds the known UTF-8 encodings to the vocabulary
		for idx = 0; idx < 256; idx++ {
			t.Vocabulary[idx] = string([]byte{byte(idx)})
		}

		if t.ShowDebugLogs {
			fmt.Println("Base vocabulary (0-255) initialized")
		}

	}

	// Text to raw bytes (initial tokens)
	tokens := []byte(text)
	intTokens := utils.ConvertSliceToInt32(tokens)

	if t.ShowDebugLogs {
		fmt.Printf("Converted text to %d initial tokens\n", len(intTokens))
		fmt.Printf("Tokens : %v\n", intTokens)
		fmt.Println("Starting merge iterations...")
	}

	for idx := startToken; idx < vocabSize; idx++ {
		if t.ShowDebugLogs {
			fmt.Printf("\n--- Iteration %d/%d ---\n", idx-startToken+1, vocabSize-startToken)
			fmt.Printf("Current token count: %d\n", len(intTokens))
		}

		// Get the most frequent pair in the raw tokens
		mostFrequentPair, numRepetitions := utils.GetMostFrequentPair(intTokens)
		if t.ShowDebugLogs {
			fmt.Printf("Most frequent pair: (%d, %d)\n", mostFrequentPair.X, mostFrequentPair.Y)
			fmt.Printf("Number of times it has been repeated : %v\n", numRepetitions)
		}

		// Mint new tokens
		newTokens := utils.MintTokens(intTokens, mostFrequentPair, idx)
		if t.ShowDebugLogs {
			fmt.Printf("After merging: %d tokens (reduced by %d)\n", len(newTokens), len(intTokens)-len(newTokens))
			fmt.Printf("After minting new token, we get : %v", newTokens)
		}

		// Updating the merges and the vocab hashmaps
		t.Merges[mostFrequentPair] = idx
		t.Vocabulary[idx] = t.Vocabulary[mostFrequentPair.X] + t.Vocabulary[mostFrequentPair.Y]
		if t.ShowDebugLogs {
			fmt.Printf("Added merge: (%d, %d) -> token %d\n", mostFrequentPair.X, mostFrequentPair.Y, idx)
		}

		intTokens = newTokens
	}
	if t.ShowDebugLogs {
		fmt.Println("\n========== TRAINING COMPLETED ==========")
		fmt.Printf("Total merges learned: %d\n", len(t.Merges))
		fmt.Printf("Final vocabulary size: %d\n", len(t.Vocabulary))
		fmt.Printf("Final token count: %d\n\n", len(intTokens))
	}
}

// Encoder
func (t *Tokenizer) Encode(text string) []int32 {
	if t.ShowDebugLogs {
		fmt.Println("========== ENCODING STARTED ==========")
		fmt.Printf("Text to encode: %q (length: %d chars)\n", text, len(text))
	}

	/*
		Step 01: Given the text, we are going to get the utf-8 encoding
		Step 02: We are going to exhaustively iterate over the sequence and search for the merges hashmap to merge new tokens
		Step 03: We will stop minting tokens once there is no addition to the existing tokens in the entire pass
	*/
	tokens := []byte(text)
	intTokens := utils.ConvertSliceToInt32(tokens)
	if t.ShowDebugLogs {
		fmt.Printf("Initial tokens: %d\n", len(intTokens))
	}

	// Loop discontinues in case there is no FRESHLY MINTED TOKEN
	var newTokens []int32
	isMintedToken := false
	idx := 0
	passNumber := 0

	for {
		passNumber++
		if t.ShowDebugLogs {
			fmt.Printf("\n--- Encoding Pass %d ---\n", passNumber)
			fmt.Printf("Starting with %d tokens\n", len(intTokens))
		}

		newTokens = []int32{} // Reset for this pass
		isMintedToken = false
		idx = 0

		for idx < len(intTokens) {
			if idx < len(intTokens)-1 {
				pair := schema.Tuple{X: intTokens[idx], Y: intTokens[idx+1]}

				// Check if we have a token trained for this specific pair
				if mergedTokenID, exists := t.Merges[pair]; exists {
					if t.ShowDebugLogs {
						fmt.Printf("  Found merge at idx %d: (%d, %d) -> %d\n", idx, pair.X, pair.Y, mergedTokenID)
					}
					isMintedToken = true
					newTokens = append(newTokens, mergedTokenID)
					idx += 2
					continue
				}
			}

			// otherwise simply add the token ID
			newTokens = append(newTokens, intTokens[idx])
			idx += 1
		}
		if t.ShowDebugLogs {
			fmt.Printf("After pass %d: %d tokens (minted: %v)\n", passNumber, len(newTokens), isMintedToken)
		}
		intTokens = newTokens

		// If no new token has been added then break
		if !isMintedToken {
			fmt.Println("No new merges found, encoding complete")
			break
		}
	}
	if t.ShowDebugLogs {
		fmt.Println("\n========== ENCODING COMPLETED ==========")
		fmt.Printf("Final encoded tokens: %d\n", len(newTokens))
		fmt.Printf("Tokens: %v\n\n", newTokens)
	}
	return newTokens
}

// Decoder
func (t *Tokenizer) Decode(tokens []int32) string {
	if t.ShowDebugLogs {
		fmt.Println("========== DECODING STARTED ==========")
		fmt.Printf("Tokens to decode: %d\n", len(tokens))
		fmt.Printf("Tokens: %v\n", tokens)
	}

	decodedString := ""

	for i, token := range tokens {
		mergedString, exists := t.Vocabulary[token]
		if !exists {
			log.Fatalf("ERROR at token index %d: No vocabulary entry for token %d", i, token)
		}

		decodedString += mergedString
	}

	if t.ShowDebugLogs {
		fmt.Println("\n========== DECODING COMPLETED ==========")
		fmt.Printf("Decoded string length: %d chars\n", len(decodedString))
		fmt.Printf("Decoded text: %q\n\n", decodedString)
		fmt.Printf("Vocabulary : %v\n", t.Vocabulary)
	}
	return decodedString
}
