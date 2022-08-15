package main

import (
	"EnigmaLorenz/pkg/lorenz"
	"fmt"
)

func main() {
	F := []byte{}
	alphabet := lorenz.NewITA2LSB()
	encoded, err := alphabet.AsciiToITA2("±±±±±±±±±±*O|ZY±_PFPQS|X±THI^_-£1|6£%6$4", true)
	fmt.Println(err)
	encoded2, _ := alphabet.AsciiToITA2("REGIMENT NUMBER 5 WILL ATTACK WEST POI", false)
	for i := 0; i < len(encoded2); i++ {
		F = append(F, encoded[i]^encoded2[i])
	}
	decoded, _ := alphabet.ITA2ToAscii(F, false)
	fmt.Println(decoded)
}
