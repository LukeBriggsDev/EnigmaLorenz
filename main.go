package main

import "fmt"

type rotor struct {
	name         string
	alphabet     []byte
	notchList    []uint
	turnoverList []uint
}

func main() {
	var I = rotor{
		name:         "I",
		alphabet:     []byte{'E', 'K', 'M', 'F', 'L', 'G', 'D', 'Q', 'V', 'Z', 'N', 'T', 'O', 'W', 'Y', 'H', 'X', 'U', 'S', 'P', 'A', 'I', 'B', 'R', 'C', 'J'},
		notchList:    []uint{14},
		turnoverList: []uint{8},
	}
	fmt.Printf("%s\n", I.name)
}
