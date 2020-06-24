package main

import (
	"fmt"
)

func main() {
	asciiBytes := []byte{80, 79, 83, 84, 32, 104, 116, 116, 112, 58, 47, 47, 108, 111, 99, 97, 108, 104, 111, 115, 116, 58, 56, 48, 56, 48, 47, 52, 50}
	for _, v := range asciiBytes {
		fmt.Printf(string(v))
	}
	fmt.Println("\n")
}
