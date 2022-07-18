package main

import (
	"EnigmaLorenz/internal/enigma"
)

func main() {
	I, II, III, _, UKW_B := enigma.GenerateRotors()
	I.CurrentPos = 6
	II.CurrentPos = 1
	III.CurrentPos = 0
	plugs := enigma.Plugboard{}
	plugs.AddPlug('R', 'B')
	plugs.AddPlug('W', 'S')
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
		Plugs:       plugs,
	}
	print(machine.Encrypt("AA"))
}
