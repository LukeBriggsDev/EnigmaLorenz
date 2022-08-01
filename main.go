package main

import (
	"EnigmaLorenz/pkg/enigma"
	"fmt"
)

func main() {
	rotorSet := enigma.GenerateRotors()
	machine := enigma.Enigma{
		LeftRotor:   rotorSet.II,
		CenterRotor: rotorSet.I,
		RightRotor:  rotorSet.III,
		Reflector:   rotorSet.UKW_B,
		Plugs:       enigma.NewPlugboard(),
	}

	machine.Plugs.AddPlug('A', 'T')
	machine.Plugs.AddPlug('B', 'P')
	machine.Plugs.AddPlug('C', 'W')
	machine.Plugs.AddPlug('D', 'N')
	machine.Plugs.AddPlug('E', 'Z')
	machine.Plugs.AddPlug('F', 'R')
	machine.Plugs.AddPlug('G', 'V')
	machine.Plugs.AddPlug('H', 'O')
	machine.Plugs.AddPlug('I', 'X')
	machine.Plugs.AddPlug('J', 'Q')
	machine.Plugs.AddPlug('K', 'S')
	machine.Plugs.AddPlug('L', 'U')
	machine.Plugs.AddPlug('M', 'Y')

	machine.LeftRotor.SetShownPos(1)
	machine.CenterRotor.SetShownPos(2)
	machine.RightRotor.SetShownPos(3)

	cipher, _ := machine.Encrypt("DONTPANIC", false)
	fmt.Println(cipher)

}
