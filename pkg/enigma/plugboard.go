package enigma

import "log"

// A Plugboard contains the state of all the mapping between letters.
type Plugboard struct {
	state map[byte]byte
}

// NewPlugboard is a constructor, returning an empty Plugboard
func NewPlugboard() Plugboard {
	return Plugboard{
		state: make(map[byte]byte),
	}
}

// AddPlug creates a connection between the two characters given to it.
//
// Errors
//
// letter 1 and letter 2 must be ASCII characters between A(65) - Z(90).
// A fatal error will occur if the characters are invalid.
func (p *Plugboard) AddPlug(letter1 byte, letter2 byte) {
	if !(validChars(string(letter1)) && validChars(string(letter2))) {
		log.Fatal("Invalid characters given to plugboard, must be A-Z")
	}
	p.state[letter1] = letter2
	p.state[letter2] = letter1
}

// RemovePlug removes the connection relating to the byte passed to it.
// If a connection does not exist with the character then no action is taken.
func (p *Plugboard) RemovePlug(letter byte) {
	val, exists := p.state[letter]
	if !exists {
		return
	}
	delete(p.state, val)
	delete(p.state, letter)
}

// Translate takes a byte and returns the result of passing that byte through the Plugboard.
// If the byte has no connection on the Plugboard then the original byte is returned.
func (p *Plugboard) Translate(letter byte) byte {
	val, exists := p.state[letter]
	if !exists {
		return letter
	}
	return val
}
