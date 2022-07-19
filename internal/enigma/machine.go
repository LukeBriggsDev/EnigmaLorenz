package enigma

type Enigma struct {
	LeftRotor   Rotor
	CenterRotor Rotor
	RightRotor  Rotor
	Reflector   Rotor
	Plugs       Plugboard
}

func (machine *Enigma) Encrypt(plaintext string) string {
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

		for _, rotor := range []Rotor{machine.RightRotor, machine.CenterRotor, machine.LeftRotor, machine.Reflector} {
			chr = rotor.Translate(chr)
		}

		for _, rotor := range []Rotor{machine.LeftRotor, machine.CenterRotor, machine.RightRotor} {
			chr = rotor.TranslateReverse(chr)
		}

		chr = chr + byte('A')

		chr = machine.Plugs.Translate(chr)

		cipher = append(cipher, chr)

	}
	return string(cipher)
}