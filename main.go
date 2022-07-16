package main

import (
	"EnigmaLorenz/internal/enigma"
	"fmt"
)

func encrypt(plaintext string, rot1 enigma.Rotor, rot2 enigma.Rotor, rot3 enigma.Rotor, ref enigma.Rotor) {
	cipher := []byte(plaintext)
	for _, chr := range cipher {
		chr = chr - byte('A')
		rot3.Rotate()
		fmt.Printf("%d\n", chr)
		chr = rot3.Translate(chr)
		fmt.Printf("%d\n", chr)
		chr = rot2.Translate(chr)
		fmt.Printf("%d\n", chr)
		chr = rot1.Translate(chr)
		fmt.Printf("%d\n", chr)
		chr = ref.Translate(chr)
		fmt.Printf("%d\n", chr)
		chr = rot1.TranslateReverse(chr)
		fmt.Printf("%d\n", chr)
		chr = rot2.TranslateReverse(chr)
		fmt.Printf("%c\n", chr)
		chr = rot3.TranslateReverse(chr)
		fmt.Printf("%d\n", chr)
		chr = chr + byte(65)
		fmt.Printf("%c\n", chr)
	}
}

func main() {
	I, II, III, _, UKW_B := enigma.GenerateRotors()
	I.CurrentPos = 12
	II.CurrentPos = 8
	III.CurrentPos = 20
	encrypt("D", I, II, III, UKW_B)
}
