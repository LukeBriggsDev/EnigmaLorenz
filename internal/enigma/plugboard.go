package enigma

import "log"

type Plugboard struct {
	state map[byte]byte
}

func NewPlugboard() Plugboard {
	return Plugboard{
		state: make(map[byte]byte),
	}
}

func (p *Plugboard) AddPlug(letter1 byte, letter2 byte) {
	if !(validChars(string(letter1)) && validChars(string(letter2))) {
		log.Fatal("Invalid characters given to plugboard, must be A-Z")
	}
	p.state[letter1] = letter2
	p.state[letter2] = letter1
}

func (p *Plugboard) RemovePlug(letter byte) {
	val, exists := p.state[letter]
	if !exists {
		return
	}
	delete(p.state, val)
	delete(p.state, letter)
}

func (p *Plugboard) Translate(letter byte) byte {
	val, exists := p.state[letter]
	if !exists {
		return letter
	}
	return val
}
