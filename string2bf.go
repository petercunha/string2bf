package main

import (
	"fmt"
)

const STR string = "Hello world!"	// String you want to convert
const COMPRESSION bool = true		// Change to true to compress output

func main() {
	if (!COMPRESSION) {
		v1_convert(STR)
	} else {
		v2_convert(STR)
	}
}

// Uses least amount of Go-code to compute brainfuck (13 lines!)
func v1_convert(s string) {
	ascii := []byte(s)
	for i := 0; i < len(ascii); i++ {
		for j := 0; j < int(ascii[i]); j++ {
			fmt.Print("+")
		}
		fmt.Print(">")
	}
	for i := 0; i < len(ascii); i++ {
		fmt.Print("<")
	}
	for i := 0; i < len(ascii); i++ {
		fmt.Print(".")
		fmt.Print(">")
	}
}

// Tries to compress computed brainfuck by using only one brainfuck memcell
// This generates a more compressed output
func v2_convert(s string) {
	ascii := []byte(s)
	for i := 0; i < len(ascii); i++ {
		if i == 0 {
			for j := 0; j < int(ascii[i]); j++ {
				fmt.Print("+")
			}
		} else if ascii[i-1] > ascii[i] {
			for j := 0; j < int(ascii[i-1])-int(ascii[i]); j++ {
				fmt.Print("-")
			}
		} else if ascii[i-1] < ascii[i] {
			for j := 0; j < int(ascii[i])-int(ascii[i-1]); j++ {
				fmt.Print("+")
			}
		}
		fmt.Print(".")
	}
}
