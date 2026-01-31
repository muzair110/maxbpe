package main

import (
	"flag"
	"fmt"

	"github.com/muzair110/maxbpe/bpe"
)

func main() {

	//text := "hello world!"
	var vocabSize int32 = 260
	var startToken int32 = 256

	var showDebugLogs bool
	var text string

	flag.BoolVar(&showDebugLogs, "verbose", false, "This is to show the debug logs for training and encoding / decoding")
	flag.StringVar(&text, "text", "hello world!!!? (안녕하세요!) lol122 😉", "This is the dataset to train the tokenizer")
	flag.Parse()

	tokenizer := &bpe.Tokenizer{ShowDebugLogs: showDebugLogs}

	// Test the encoder and the decoder (without training)
	tokenizer.Train(vocabSize, startToken, text)
	tokens := tokenizer.Encode(text)
	decodedText := tokenizer.Decode(tokens)

	fmt.Printf("\nOriginal Text : %s... Decoded Text : %s", text, decodedText)
	fmt.Printf("\n Is the original and the decoded text the same? : %v", text == decodedText)
}
