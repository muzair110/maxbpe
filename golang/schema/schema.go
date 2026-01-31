package schema

// Schema for Tuples to collect the byte pairs
type Tuple struct {
	X, Y int32
}

// This struct contains the merges and the vocabulary hashmaps
// Very useful for encoding and decoding later on
type Assets struct {
	Merges     map[Tuple]int32
	Vocabulary map[int32]string
}

// The interface defining the training, encoding and decoding methods
type BpeTokenizer interface {
	// This takes in the following arguments
	// 		vocabSize : int
	//		startToken : int
	// 		raw string
	Train(int32, int32, string)

	// For encoding
	Encode(string) []int32

	// For decoding
	Decode([]int32) string
}
