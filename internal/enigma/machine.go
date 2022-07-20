package enigma

import (
	"errors"
)

type Enigma struct {
	LeftRotor   Rotor
	CenterRotor Rotor
	RightRotor  Rotor
	FourthRotor Rotor
	Reflector   Rotor
	Plugs       Plugboard
}

func validChars(text string) bool {
	for _, chr := range text {
		if chr < 'A' || chr > 'Z' {
			return false
		}
	}
	return true
}

func (machine *Enigma) Encrypt(plaintext string, useFourthRotor bool) (string, error) {
	if !validChars(plaintext) {
		return "", errors.New("enigma input must be capitalized ascii letters only")
	}

	var cipher []byte
	for _, chr := range []byte(plaintext) {
		chr = machine.Plugs.Translate(chr)
		chr = chr - byte('A')
		if machine.CenterRotor.AtNotch() {
			machine.LeftRotor.Rotate()
			// Double stepping of center rotor
			machine.CenterRotor.Rotate()
		}
		if machine.RightRotor.AtNotch() {
			machine.CenterRotor.Rotate()
		}
		machine.RightRotor.Rotate()

		path := []Rotor{machine.RightRotor, machine.CenterRotor, machine.LeftRotor}
		if useFourthRotor {
			path = append(path, machine.FourthRotor)
		}

		for _, rotor := range path {
			chr = rotor.Translate(chr)
		}

		chr = machine.Reflector.Translate(chr)

		for rotorIndex := len(path) - 1; rotorIndex >= 0; rotorIndex-- {
			chr = path[rotorIndex].TranslateReverse(chr)
		}

		chr = chr + byte('A')

		chr = machine.Plugs.Translate(chr)

		cipher = append(cipher, chr)

	}
	return string(cipher), nil
}
