package main

import (
	"fmt"
	"log"
)

func main() {
	//ctx := context.Background()
	text := "This is some text dataset hello, and hi some words!"

	tokens, error := characterEncoding(text)
	if error != nil {
		log.Fatalf("Error creating naive tokens : %v", error)
	}
	fmt.Printf("Tokens : %v", tokens)

}
